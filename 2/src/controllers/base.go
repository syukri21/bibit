package controllers

import (
	"bibit-test/src/models/proto"
	"github.com/sarulabs/di"
)

type Controller struct {
	MovieHttp MovieHTTP
	MovieGRPC proto.MoviesServer
}

func NewController(ioc di.Container) *Controller {
	return &Controller{
		MovieHttp: NewMovieHTTP(ioc),
		MovieGRPC: NewMovieGRPC(ioc),
	}
}
