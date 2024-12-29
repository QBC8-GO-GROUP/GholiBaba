package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/QBC8-GO-GROUP/GholiBaba/api/handlers/http"
	"github.com/QBC8-GO-GROUP/GholiBaba/app"
	"github.com/QBC8-GO-GROUP/GholiBaba/config"
)

var configPath = flag.String("config", "config.json", "service configuration file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	fmt.Println("c:", c)

	appContainer := app.MustNewApp(c)

	log.Fatal(http.Run(appContainer, c.Server))
}
