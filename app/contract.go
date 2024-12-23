package app

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/config"
	"gorm.io/gorm"

	companyPort "github.com/QBC8-GO-GROUP/GholiBaba/internal/company/port"
)

type App interface {
	DB() *gorm.DB

	Config(ctx context.Context) config.Config
	CompanyService(ctx context.Context) companyPort.Service
}
