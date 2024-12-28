package app

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/config"
	"gorm.io/gorm"

	travelPort "github.com/QBC8-GO-GROUP/GholiBaba/internal/travel/port"
)

type App interface {
	DB() *gorm.DB

	Config(ctx context.Context) config.Config
	TravelService(ctx context.Context) travelPort.Service
}
