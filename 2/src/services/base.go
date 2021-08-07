package services

import "github.com/sarulabs/di"

type Service struct {
	Movie Movie
}

func NewService(ioc di.Container) *Service {
	return &Service{
		Movie: NewMovie(ioc),
	}
}
