package app

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/config"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel"
	travelPort "github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/port"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/postgres"

	appCtx "github.com/QBC8-GO-GROUP/GholiBaba/pkg/context"
	"gorm.io/gorm"
)

type app struct {
	db            *gorm.DB
	cfg           config.Config
	travelService travelPort.Service
}

func (a *app) DB() *gorm.DB {
	return a.db
}
func (a *app) TravelService(ctx context.Context) travelPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.travelService == nil {
			a.travelService = a.travelServiceWithDB(a.db)
		}
		return a.travelService
	}

	return a.travelServiceWithDB(db)
}

func (a *app) travelServiceWithDB(db *gorm.DB) travelPort.Service {
	return travel.NewService(storage.NewTravelRepo(db))
}

func (a *app) Config(ctx context.Context) config.Config {
	return a.cfg
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Password,
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		DBName: a.cfg.DB.Database,
		Schema: a.cfg.DB.Schema,
	})

	postgres.GormMigrations(db)

	if err != nil {
		return err
	}

	a.db = db

	return nil
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	a.travelService = travel.NewService(storage.NewTravelRepo(a.db))
	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
