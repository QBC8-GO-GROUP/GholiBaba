package app

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/config"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/adapters/storage"

	"github.com/QBC8-GO-GROUP/GholiBaba/internal/user"
	"github.com/QBC8-GO-GROUP/GholiBaba/internal/user/domain"
	userPort "github.com/QBC8-GO-GROUP/GholiBaba/internal/user/port"

	appCtx "github.com/QBC8-GO-GROUP/GholiBaba/pkg/context"
	"github.com/QBC8-GO-GROUP/GholiBaba/pkg/postgres"

	"gorm.io/gorm"
)

type app struct {
	db          *gorm.DB
	cfg         config.Config
	userService userPort.Service
}

// CodeVerificationService implements App.

func (a *app) DB() *gorm.DB {
	return a.db
}
func (a *app) UserService(ctx context.Context) userPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.userService == nil {
			a.userService = a.userServiceWithDB(a.db)
		}
		return a.userService
	}

	return a.userServiceWithDB(db)
}

func (a *app) userServiceWithDB(db *gorm.DB) userPort.Service {
	return user.NewService(storage.NewUserRepo(db))
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

	app.userService = user.NewService(storage.NewUserRepo(app.db))
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
		{Policy: domain.PolicyUnknown, Resource: "/api/v1/user", Scope: "create"},
		{Policy: domain.PolicyUnknown, Resource: "/api/v1/user/update", Scope: "update"},
		{Policy: domain.PolicyUnknown, Resource: "/api/v1/user/:id", Scope: "delete"},
		{Policy: domain.PolicyUnknown, Resource: "/api/v1/user/:id", Scope: "read"},
	}
	return permissions
}
