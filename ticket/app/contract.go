package app

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/config"
	"gorm.io/gorm"

	ticketPort "github.com/QBC8-GO-GROUP/GholiBaba/ticket/internal/ticket/port"
)

type App interface {
	DB() *gorm.DB

	Config(ctx context.Context) config.Config
	TicketService(ctx context.Context) ticketPort.Service
}
