package services

import (
	"bibit-test/src/constants"
	"bibit-test/src/models"
	"bibit-test/src/repositories"
	"github.com/sarulabs/di"
)

type Movie interface {
	Search(params models.MoviesSearchParams) (result *models.MoviesSearchResponse, err error)
	GetDetail(params models.MoviesGetDetailParams) (result *models.MovieDetailResponse, err error)
	GetDetailMany(params map[string][]string) (result *[]models.MovieDetailResponse, err error)
}

type MovieImpl struct {
	repository *repositories.Repository
}

func (m *MovieImpl) Search(params models.MoviesSearchParams) (result *models.MoviesSearchResponse, err error) {
	result, err = m.repository.Movie.Search(params)
	return
}

func (m *MovieImpl) GetDetail(params models.MoviesGetDetailParams) (result *models.MovieDetailResponse, err error) {
	result, err = m.repository.Movie.GetDetail(params)
	return
}

func (m *MovieImpl) GetDetailMany(params map[string][]string) (result *[]models.MovieDetailResponse, err error) {
	panic("implement me")
}

func NewMovie(ioc di.Container) Movie {
	return &MovieImpl{
		repository: ioc.Get(constants.REPOSITORY).(*repositories.Repository),
	}
}
