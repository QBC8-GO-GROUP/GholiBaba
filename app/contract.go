package app

import (
	"github.com/QBC8-GO-GROUP/GholiBaba/config"
	userPort "github.com/QBC8-GO-GROUP/GholiBaba/internal/user/port"
)

type App interface {
	UserService() userPort.Service
	Config() config.Config
}
