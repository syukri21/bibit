package repositories

import (
	"bibit-test/src/constants"
	"bibit-test/src/models/model"
	"bibit-test/src/utils/proxy"
	"encoding/json"
	"github.com/sarulabs/di"
	"strconv"
)

type Movie interface {
	Search(params model.MoviesSearchParams) (*model.MoviesSearchResponse, error)
	GetDetail(params model.MoviesGetDetailParams) (*model.MovieDetailResponse, error)
}

type MovieImpl struct {
	omdbapi *proxy.Proxy
}

func (m MovieImpl) GetDetail(params model.MoviesGetDetailParams) (result *model.MovieDetailResponse, err error) {
	bytes, err := m.omdbapi.Query(map[string]string{
		"i": params.ID,
	}).Get()

	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}

	return
}

func (m MovieImpl) Search(params model.MoviesSearchParams) (result *model.MoviesSearchResponse, err error) {
	bytes, err := m.omdbapi.Query(map[string]string{
		"s":    params.Search,
		"page": strconv.Itoa(params.Page),
	}).Get()

	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}

	return
}

func NewMovieRepository(ioc di.Container) Movie {
	return &MovieImpl{
		omdbapi: ioc.Get(constants.PROXY_OMDBAPI).(*proxy.Proxy),
	}
}
