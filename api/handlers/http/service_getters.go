package http

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/api/service"
	"github.com/QBC8-GO-GROUP/GholiBaba/app"
	"github.com/QBC8-GO-GROUP/GholiBaba/config"
)

// company service transient instance handler
func companyServiceGetter(appContainer app.App, cfg config.ServerConfig) ServiceGetter[*service.CompanyService] {
	return func(ctx context.Context) *service.CompanyService {
		return service.NewCompanyService(appContainer.CompanyService(ctx),
			cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute, appContainer.CompanyService(ctx))
	}
}
