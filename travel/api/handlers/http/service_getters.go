package http

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/api/service"
	"github.com/QBC8-GO-GROUP/GholiBaba/app"
	"github.com/QBC8-GO-GROUP/GholiBaba/config"
)

// travel service transient instance handler
func travelServiceGetter(appContainer app.App, cfg config.ServerConfig) ServiceGetter[*service.TravelService] {
	return func(ctx context.Context) *service.TravelService {
		return service.NewTravelService(appContainer.TravelService(ctx))
	}
}
