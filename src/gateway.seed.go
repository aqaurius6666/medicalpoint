package main

import (
	"context"
	"log"

	"github.com/medicalpoint/gateway/src/db"
	"github.com/medicalpoint/gateway/src/db/interface/user"
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
	err = seedSuperAdmin(service, appCtx.String("super-admin-id"))
	if err != nil {
		log.Println("seed users fail, err:", err)
	}

	return nil
}
func seedSuperAdmin(db db.GateWayServiceRepo, id string) error {
	encryptedKey := "4+H3nDRdP0xi6e3/Yo94mO0AmMcIgLMZuV/TLIWvS6fBbS7dEBtIK2h5OQOE/JpUcTnnBLnHVZRscdb4"
	_, err := db.CreateUser(&user.User{
		UserID:              &id,
		EncryptedPrivateKey: &encryptedKey,
	})
	if err != nil {
		return err
	}
	log.Println("seed user successfully")
	return nil
}
