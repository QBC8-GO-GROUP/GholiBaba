package app

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/config"
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/pkg/broadcast"
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/pkg/nats"
)

type app struct {
	nats            *nats.Connection
	broadcastServer broadcast.Server
	cfg             config.NatsConfig
}

func NewApp(cfg config.NatsConfig, nats *nats.Connection, broadcastSrv broadcast.Server) App {
	return &app{
		cfg:             cfg,
		nats:            nats,
		broadcastServer: broadcastSrv,
	}
}

func (a *app) Nats() *nats.Connection {
	return a.nats
}

func (a *app) Config() config.NatsConfig {
	return a.cfg
}

func (a *app) BroadCast() broadcast.Server {
	return a.broadcastServer
}
