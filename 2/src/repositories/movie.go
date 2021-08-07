package repositories

import (
	"bibit-test/src/constants"
	"bibit-test/src/models"
	"bibit-test/src/utils/proxy"
	"encoding/json"
	"github.com/sarulabs/di"
	"strconv"
)

type Movie interface {
	Search(params models.MoviesSearchParams) (*models.MoviesSearchResponse, error)
	GetDetail(params models.MoviesGetDetailParams) (*models.MovieDetailResponse, error)
}

type MovieImpl struct {
	omdbapi *proxy.Proxy
}

func (m MovieImpl) GetDetail(params models.MoviesGetDetailParams) (result *models.MovieDetailResponse, err error) {
	bytes, err := m.omdbapi.Query(map[string]string{
		"i": params.ID,
	}).Get()

	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}

	return
}

func (m MovieImpl) Search(params models.MoviesSearchParams) (result *models.MoviesSearchResponse, err error) {
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
