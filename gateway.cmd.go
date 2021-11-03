package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sonntuet1997/medical-chain-utils/common"
	"github.com/urfave/cli/v2"
)

const (
	serviceName = "gateway"
)

var logger *logrus.Logger

func main() {
	logger = logrus.New()
	if err := makeApp().Run(os.Args); err != nil {
		logger.WithField("err", err).Error("shutting down due to error")
		_ = os.Stderr.Sync()
		os.Exit(1)
	}
}

func makeApp() *cli.App {
	app := &cli.App{
		Name:                 serviceName,
		Version:              "v1.0.1",
		EnableBashCompletion: true,
		Compiled:             time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Vu Nguyen",
				Email: "aqaurius6666@gmail.com",
			},
		},
		Copyright: "(c) 2021 SOTANEXT inc.",
		Action:    runMain,
		Commands: []*cli.Command{
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Usage:   "run server",
				Action:  runMain,
			},
			{
				Name:    "seed-data",
				Aliases: []string{"sd"},
				Usage:   "seed data",
				Action:  seedData,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "clean",
						EnvVars: []string{"CLEAN_DB"},
						Usage:   "Clean DB before seeding",
					},
				},
			},
			{
				Name:    "clean",
				Aliases: []string{"c"},
				Usage:   "clean DB",
				Action:  clean,
			},
		},
		Flags: append([]cli.Flag{
			&cli.StringFlag{
				Name:     "db-uri",
				Required: true,
				EnvVars:  []string{"DB_URI"},
				Usage:    "The URI for connecting to database (supported URIs: in-memory://, postgresql://auth@host:26257/linkgraph?sslmode=disable)",
			},
			&cli.StringFlag{
				Name:    "cosmos-endpoint",
				Value:   "localhost:9090",
				EnvVars: []string{"COSMOS_ENDPOINT"},
				Usage:   "Cosmos GRPC endpoint",
			},
			&cli.StringFlag{
				Name:     "mnemonic",
				Value:    "",
				Required: true,
				EnvVars:  []string{"MNEMONIC"},
				Usage:    "Mnemonic of blockchain admin account",
			},
			&cli.StringFlag{
				Name:    "chain-id",
				Value:   "medichain",
				EnvVars: []string{"CHAIN_ID"},
				Usage:   "Cosmos blockchain id",
			},
		},
			append(common.CommonGRPCFlag,
				common.LoggerFlag...)...),
	}
	return app

}
