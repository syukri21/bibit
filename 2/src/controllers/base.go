package controllers

import (
	"github.com/sarulabs/di"
)

type Controller struct {
	MovieHttp MovieHTTP
	MovieGRPC MovieGRPC
}

func NewController(ioc di.Container) *Controller {
	return &Controller{
		MovieHttp: NewMovieHTTP(ioc),
		MovieGRPC: NewMovieGRPC(ioc),
	}
}
