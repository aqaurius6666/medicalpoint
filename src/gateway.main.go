// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"cloud.google.com/go/profiler"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/medicalpoint/gateway/src/api"
	"github.com/medicalpoint/gateway/src/db"
	"github.com/medicalpoint/gateway/src/services/cosmos"
	"github.com/medicalpoint/gateway/src/services/swagger"
	"github.com/sonntuet1997/medical-chain-utils/common_service"
	pb2 "github.com/sonntuet1997/medical-chain-utils/common_service/pb"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/tendermint/spm/cosmoscmd"
	"github.com/urfave/cli/v2"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func runMain(appCtx *cli.Context) error {

	var wg sync.WaitGroup
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()
	cosmoscmd.SetPrefixes("medipoint")
	server, err := InitGateWayServer(
		ctx,
		AppOptions{
			cliContext: appCtx,
			dbDsn:      db.DBDsn(appCtx.String("db-uri")),
			ChainId:    cosmos.ChainID(appCtx.String("chain-id")),
			KeyRing:    keyring.NewUnsafe(keyring.NewInMemory()),
			CosmosEp:   cosmos.CosmosEndpoint(appCtx.String("cosmos-endpoint")),
			Mne:        cosmos.Mnemonic(appCtx.String("mnemonic")),
			Key:        appCtx.String("encrypt-key"),
		},
	)
	if err != nil {
		return err
	}
	defer func(mainServer *api.GateWayServer) {
		err := mainServer.Repo.Close()
		if err != nil {
			panic("error when closing DB Connection!")
		}
	}(server)
	err = server.Repo.Migrate()
	if err != nil {
		return err
	}
	if appCtx.Bool("disable-tracing") {
		logger.Info("Tracing disabled.")
	} else {
		logger.Info("Tracing enabled.")
		go initTracing()
	}
	if appCtx.Bool("disable-profiler") {
		logger.Info("Profiling disabled.")
	} else {
		logger.Info("Profiling enabled.")
		go initProfiling(serviceName, appCtx.String("runtime-version"))
	}

	// Start gRPC server
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", appCtx.Int("grpc-port")))
	if err != nil {
		return err
	}
	defer func() { _ = grpcListener.Close() }()
	var srv *grpc.Server
	commonServer := common_service.NewCommonServiceServer(logger, appCtx.Bool("allow-kill"))
	wg.Add(1)
	go func() {
		defer wg.Done()
		if appCtx.Bool("disable-stats") {
			logger.Info("Stats disabled.")
			srv = grpc.NewServer()
		} else {
			logger.Info("Stats enabled.")
			srv = grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
		}
		healthpb.RegisterHealthServer(srv, commonServer)
		pb2.RegisterCommonServiceServer(srv, commonServer)
		reflection.Register(srv)
		logger.WithField("port", appCtx.Int("grpc-port")).Info("listening for gRPC connections")
		if err := srv.Serve(grpcListener); err != nil {
			logger.Fatalf("failed to serve: %v", err)
		}
	}()
	{
		// _, err = cosmosServiceClient.AddAccountFromMnemonic("admin", appCtx.String("mnemonic"))
		// if err != nil {
		// logger.Errorf("error while adding admin mnemonic: %v", err)
		// }
	}

	apiSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", appCtx.Int("http-port")),
		Handler: server.G,
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err != nil {
			logger.Errorf("err: %v", err)
		}
		server.RegisterEndpoint()
		server.G.GET("/swagger/*any", swagger.CustomWrapHandler(
			&ginSwagger.Config{
				URL:         "gateway.json",
				DeepLinking: true,
			},
			swaggerFiles.Handler))
		logger.WithField("port", appCtx.Int("http-port")).Info("listening for api connections")
		if err := apiSrv.ListenAndServe(); err != nil {
			logger.Printf("listen: %s\n", err)
		}
	}()
	// Start pprof server
	pprofListener, err := net.Listen("tcp", fmt.Sprintf(":%d", appCtx.Int("pprof-port")))
	if err != nil {
		return err
	}
	defer func() { _ = pprofListener.Close() }()

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.WithField("port", appCtx.Int("pprof-port")).Info("listening for pprof requests")
		sSrv := new(http.Server)
		_ = sSrv.Serve(pprofListener)
	}()
	// Start signal watcher
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGKILL)
		select {
		case s := <-sigCh:
			logger.WithField("signal", s.String()).Infof("shutting down due to signal")
			srv.Stop()
			_ = grpcListener.Close()
			_ = pprofListener.Close()
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := apiSrv.Shutdown(ctx); err != nil {
				logger.Fatal("Server Shutdown:", err)
			}
			cancelFn()
			logger.WithField("signal", s.String()).Infof("shutdown success!")
		case <-ctx.Done():
		}
	}()
	// Keep running until we receive a signal
	wg.Wait()
	return nil
}

func initTracing() {
	// initJaegerTracing()
	initStackdriverTracing()
}

func initStats(exporter *stackdriver.Exporter) {
	view.SetReportingPeriod(60 * time.Second)
	view.RegisterExporter(exporter)
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		logger.Warn("Error registering default server views")
	} else {
		logger.Info("Registered default server views")
	}
}

func initStackdriverTracing() {
	// TODO(ahmetb) this method is duplicated in other microservices using Go
	// since they are not sharing packages.
	for i := 1; i <= 3; i++ {
		exporter, err := stackdriver.NewExporter(stackdriver.Options{})
		if err != nil {
			logger.Infof("failed to initialize stackdriver exporter: %+v", err)
		} else {
			trace.RegisterExporter(exporter)
			logger.Info("registered Stackdriver tracing")

			// Register the views to collect server stats.
			initStats(exporter)
			return
		}
		d := time.Second * 10 * time.Duration(i)
		logger.Infof("sleeping %v to retry initializing Stackdriver exporter", d)
		time.Sleep(d)
	}
	logger.Warn("could not initialize Stackdriver exporter after retrying, giving up")
}

func initProfiling(service, version string) {
	// TODO(ahmetb) this method is duplicated in other microservices using Go
	// since they are not sharing packages.
	for i := 1; i <= 3; i++ {
		if err := profiler.Start(profiler.Config{
			Service:        service,
			ServiceVersion: version,
			// ProjectID must be set if not running on GCP.
			// ProjectID: "my-project",
		}); err != nil {
			logger.Warnf("failed to start profiler: %+v", err)
		} else {
			logger.Info("started Stackdriver profiler")
			return
		}
		d := time.Second * 10 * time.Duration(i)
		logger.Infof("sleeping %v to retry initializing Stackdriver profiler", d)
		time.Sleep(d)
	}
	logger.Warn("could not initialize Stackdriver profiler after retrying, giving up")
}
