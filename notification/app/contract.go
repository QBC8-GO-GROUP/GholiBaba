package app

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/config"
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/pkg/broadcast"
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/pkg/nats"
)

type App interface {
	Nats() *nats.Connection
	Config() config.NatsConfig
	BroadCast() broadcast.Server
}
