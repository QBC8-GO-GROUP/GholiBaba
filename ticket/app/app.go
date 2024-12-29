package app

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/config"

	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/internal/ticket"
	ticketPort "github.com/QBC8-GO-GROUP/GholiBaba/ticket/internal/ticket/port"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/pkg/adapters/storage"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/pkg/postgres"

	appCtx "github.com/QBC8-GO-GROUP/GholiBaba/ticket/pkg/context"
	"gorm.io/gorm"
)

type app struct {
	db            *gorm.DB
	cfg           config.Config
	ticketService ticketPort.Service
}

func (a *app) DB() *gorm.DB {
	return a.db
}
func (a *app) TicketService(ctx context.Context) ticketPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.ticketService == nil {
			a.ticketService = a.ticketServiceWithDB(a.db)
		}
		return a.ticketService
	}

	return a.ticketServiceWithDB(db)
}

func (a *app) ticketServiceWithDB(db *gorm.DB) ticketPort.Service {
	return ticket.NewService(storage.NewTicketRepo(db))
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

	a.ticketService = ticket.NewService(storage.NewTicketRepo(a.db))
	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
