package services

import (
	"bibit-test/src/constants"
	"bibit-test/src/models/model"
	"bibit-test/src/repositories"
	"encoding/json"
	"errors"
	"github.com/sarulabs/di"
	"net/http"
)

type Movie interface {
	Search(params model.MoviesSearchParams) (result *model.MoviesSearchResponse, code int, err error)
	GetDetail(params model.MoviesGetDetailParams) (result *model.MovieDetailResponse, code int, err error)
	GetDetailMany(params map[string][]string) (result *[]model.MovieDetailResponse, code int, err error)
}

type MovieImpl struct {
	repository *repositories.Repository
}

func (m *MovieImpl) Search(params model.MoviesSearchParams) (result *model.MoviesSearchResponse, code int, err error) {
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

	go func() {
		data, _ := json.Marshal(result)
		_ = m.repository.MovieLog.CreateSearchLog(&model.MovieSearchLog{
			Data: data,
		})
	}()

	return
}

func (m *MovieImpl) GetDetail(params model.MoviesGetDetailParams) (result *model.MovieDetailResponse, code int, err error) {
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

	defer func() {
		data, _ := json.Marshal(result)
		_ = m.repository.MovieLog.CreateGetDetailLog(&model.MovieGetDetailLog{
			Data: data,
		})
	}()

	return
}

func (m *MovieImpl) GetDetailMany(map[string][]string) (result *[]model.MovieDetailResponse, code int, err error) {
	panic("implement me")
}

func NewMovie(ioc di.Container) Movie {
	return &MovieImpl{
		repository: ioc.Get(constants.REPOSITORY).(*repositories.Repository),
	}
}
