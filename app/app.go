package app

import (
	"context"
	"subscription-aggregator-api/config"
	"subscription-aggregator-api/db"
	"subscription-aggregator-api/manager"
	"subscription-aggregator-api/server"
	"subscription-aggregator-api/storage"
)

func MustStart() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.MustLoad()
	if err != nil {
		return err
	}

	dbManager := db.NewDBManager()

	if err := dbManager.InitDB(cfg.DBCfg); err != nil {
		return err
	}
	defer dbManager.Close()

	sqlStorage := storage.NewSQL(dbManager.DB)

	subscriptionManager := manager.New(sqlStorage)

	server := server.Init(ctx, subscriptionManager)

	return server.MustRun(cfg.SrvCfg)
}


