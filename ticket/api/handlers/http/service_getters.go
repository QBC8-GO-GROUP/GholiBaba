package http

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/api/service"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/app"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/config"
)

// ticket service transient instance handler
func ticketServiceGetter(appContainer app.App, cfg config.ServerConfig) ServiceGetter[*service.TicketService] {
	return func(ctx context.Context) *service.TicketService {
		return service.NewTicketService(appContainer.TicketService(ctx))
	}
}
