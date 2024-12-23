package app

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/config"
	"gorm.io/gorm"

	userPort "github.com/QBC8-GO-GROUP/GholiBaba/internal/user/port"
)

type App interface {
	DB() *gorm.DB

	Config(ctx context.Context) config.Config
	UserService(ctx context.Context) userPort.Service
}
