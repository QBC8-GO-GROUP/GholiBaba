package main

import (
	"context"
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/api/handlers/http"
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/app"
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/config"
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/pkg/broadcast"
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/pkg/nats"
)

func main() {

	cfg := config.MustReadConfig("config.json")

	natsConn := nats.NewConnection(cfg.Host, cfg.Port, cfg.Subject)
	natsConn.MustConnect()

	ctx := context.Background()
	broadcastServer := broadcast.NewBroadcastServer(ctx, natsConn.Ch)

	container := app.NewApp(cfg, natsConn, broadcastServer)

	http.Run(container)

}
