package app

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/payment/config"
	myPostgres "github.com/QBC8-GO-GROUP/GholiBaba/payment/pkg/postgres"
	"gorm.io/gorm"
)

type App struct {
	config config.Config
	db     *gorm.DB
}

func NewApp(cfg config.Config) (*App, error) {
	app := &App{
		config: cfg,
	}

	err := app.setDB()

	if err != nil {
		return nil, err
	}

	return app, nil
}

func MustNewApp(cfg config.Config) *App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}

func (app *App) setDB() error {
	var cfg = app.config.DB
	db, err := myPostgres.NewPsqlGormConnection(myPostgres.DBConnOptions{
		User:   cfg.User,
		Pass:   cfg.Password,
		Host:   cfg.Host,
		Port:   cfg.Port,
		DBName: cfg.Database,
		Schema: cfg.Schema,
	})

	if err != nil {
		return err
	}

	app.db = db
	return nil

}
