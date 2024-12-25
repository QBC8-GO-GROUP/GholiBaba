package main

import (
	"fmt"
	"github.com/QBC8-GO-GROUP/GholiBaba/nofitication/config"
)

func main() {

	cfg := config.MustReadConfig("config.json")

	fmt.Println(cfg)

}
