package repositories

import (
	"bibit-test/src/constants"
	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type Repository struct {
	Movie    Movie
	MovieLog MovieLog
}

func NewRepository(ioc di.Container) *Repository {
	return &Repository{
		Movie:    NewMovieRepository(ioc),
		MovieLog: NewMovieLog(ioc),
	}
}

func getDatabase(ioc di.Container) *gorm.DB {
	return ioc.Get(constants.DATABASE).(*gorm.DB)
}
