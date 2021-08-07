package services

import (
	"bibit-test/src/constants"
	"bibit-test/src/models"
	"bibit-test/src/repositories"
	"errors"
	"github.com/sarulabs/di"
	"net/http"
)

type Movie interface {
	Search(params models.MoviesSearchParams) (result *models.MoviesSearchResponse, code int, err error)
	GetDetail(params models.MoviesGetDetailParams) (result *models.MovieDetailResponse, code int, err error)
	GetDetailMany(params map[string][]string) (result *[]models.MovieDetailResponse, code int, err error)
}

type MovieImpl struct {
	repository *repositories.Repository
}

func (m *MovieImpl) Search(params models.MoviesSearchParams) (result *models.MoviesSearchResponse, code int, err error) {
	result, err = m.repository.Movie.Search(params)
	if err != nil {
		code = http.StatusBadGateway
		return
	}
	if !result.Response {
		err = errors.New("Not Found!!")
		code = http.StatusNotFound
		return
	}
	return
}

func (m *MovieImpl) GetDetail(params models.MoviesGetDetailParams) (result *models.MovieDetailResponse, code int, err error) {
	result, err = m.repository.Movie.GetDetail(params)
	if err != nil {
		code = http.StatusBadGateway
		return
	}
	if !result.Response {
		err = errors.New("Not Found!!")
		code = http.StatusNotFound
		return
	}
	return
}

func (m *MovieImpl) GetDetailMany(params map[string][]string) (result *[]models.MovieDetailResponse, code int, err error) {
	panic("implement me")
}

func NewMovie(ioc di.Container) Movie {
	return &MovieImpl{
		repository: ioc.Get(constants.REPOSITORY).(*repositories.Repository),
	}
}
