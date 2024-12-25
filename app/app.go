package app

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/config"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/domain"
	travelPort "github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/port"

	appCtx "github.com/QBC8-GO-GROUP/GholiBaba/pkg/context"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/postgres"

	"gorm.io/gorm"
)

type app struct {
	db            *gorm.DB
	cfg           config.Config
	travelService travelPort.Service
}

// CodeVerificationService implements App.

func (a *app) DB() *gorm.DB {
	return a.db
}
func (a *app) CompanyService(ctx context.Context) travelPort.Service {
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
		Company: a.cfg.DB.Company,
		Pass:    a.cfg.DB.Password,
		Host:    a.cfg.DB.Host,
		Port:    a.cfg.DB.Port,
		DBName:  a.cfg.DB.Database,
		Schema:  a.cfg.DB.Schema,
	})

	postgres.GormMigrations(db)

	if err != nil {
		return err
	}

	a.db = db

	if err != nil {
		return err
	}

	return nil
}

func NewApp(cfg config.Config) (App, error) {
	app := &app{
		cfg: cfg,
	}

	if err := app.setDB(); err != nil {
		return nil, err
	}

	app.travelService = travel.NewService(storage.NewCompanyRepo(app.db))
	return app, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}

func generatePermissions() []domain.Permission {
	permissions := []domain.Permission{
		{Policy: domain.PolicyUnknown, Resource: "/api/v1/travel", Scope: "create"},
		{Policy: domain.PolicyUnknown, Resource: "/api/v1/travel/update", Scope: "update"},
		{Policy: domain.PolicyUnknown, Resource: "/api/v1/travel/:id", Scope: "delete"},
		{Policy: domain.PolicyUnknown, Resource: "/api/v1/travel/:id", Scope: "read"},
	}
	return permissions
}
