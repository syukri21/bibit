package repositories

import "github.com/sarulabs/di"

type Repository struct {
	Movie Movie
}

func NewRepository(ioc di.Container) *Repository {
	return &Repository{
		Movie: NewMovieRepository(ioc),
	}
}
