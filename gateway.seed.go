package main

import (
	"context"
	"log"

	"github.com/medicalpoint/gateway/src/db"
	"github.com/medicalpoint/gateway/src/db/cockroach/user"
	"github.com/sonntuet1997/medical-chain-utils/common"
	"github.com/urfave/cli/v2"
)

func clean(appCtx *cli.Context) error {
	log := common.InitLogger(appCtx)
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()
	service, err := db.InitGateWayServiceRepo(ctx, log, db.DBDsn(appCtx.String("db-uri")))
	if err != nil {
		return err
	}
	defer func(db db.GateWayServiceRepo) {
		e := db.Close()
		if e != nil {
			panic("cannot close DB")
		}
	}(service)
	if err != nil {
		return err
	}
	return service.Drop()
}

func seedData(appCtx *cli.Context) error {
	log := common.InitLogger(appCtx)
	log.Info("Seed starting!")
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	service, err := db.InitGateWayServiceRepo(ctx, log, db.DBDsn(appCtx.String("db-uri")))
	if err != nil {
		return err
	}
	defer func(db db.GateWayServiceRepo) {
		e := db.Close()
		if e != nil {
			panic("cannot close DB")
		}
	}(service)
	if appCtx.Bool("clean") {
		log.Debugf("start cleaning DB")
		err := clean(appCtx)
		if err != nil {
			log.Errorf("failed cleaning DB")
			return err
		}
		log.Debugf("sucessed cleaning DB")
	}
	err = service.Migrate()
	if err != nil {
		return err
	}
	err = seedUser(service)
	if err != nil {
		log.Println("seed users fail, err:", err)
	}

	return nil
}

func seedUser(db db.GateWayServiceRepo) error {
	err := db.RawSql(user.UserSeedDataCDB)
	if err != nil {
		return err
	}
	log.Println("seed user successfully")
	return nil
}
