package main

import (
	"bibit-test/src/controllers"
	"bibit-test/src/route/grpc"
	"bibit-test/src/route/http"
	"github.com/sarulabs/di"
)

type App struct {
	http       *http.HTTP
	controller *controllers.Controller
	grpc       *grpc.GRPC
}

func NewApp(ioc di.Container) App {
	httpApp := http.New(ioc)
	grpcApp := grpc.New(ioc)

	return App{
		http: httpApp,
		grpc: grpcApp,
	}
}
