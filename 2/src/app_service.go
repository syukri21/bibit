package main

import (
	"bibit-test/src/controllers"
	"bibit-test/src/route/http"
	"github.com/sarulabs/di"
)

type App struct {
	http       *http.Route
	controller *controllers.Controller
}

func NewApp(ioc di.Container) App {
	return App{
		http: http.NewRoute(ioc),
	}
}
