package controllers

import "github.com/sarulabs/di"

type MovieGRPC interface {
}

type MovieGRPCImpl struct {
}

func NewMovieGRPC(ioc di.Container) MovieGRPC {
	return &MovieGRPCImpl{}
}
