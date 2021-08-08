package repositories

import "bibit-test/src/models/model"

type MovieMock struct {
	SearchData    *model.MoviesSearchResponse
	GetDetailData *model.MovieDetailResponse
}

func (m MovieMock) Search(model.MoviesSearchParams) (data *model.MoviesSearchResponse, err error) {
	data = m.SearchData
	return
}

func (m MovieMock) GetDetail(model.MoviesGetDetailParams) (data *model.MovieDetailResponse, err error) {
	data = m.GetDetailData
	return
}

func NewMovieMock(data MovieMock) Movie {
	return &data
}

type MovieLogMock struct {
	LogCreateError error
}

func (m MovieLogMock) CreateSearchLog(*model.MovieSearchLog) (err error) {
	return
}

func (m MovieLogMock) CreateGetDetailLog(*model.MovieGetDetailLog) (err error) {
	return
}

func NewMovieLogMock(data MovieLogMock) MovieLog {
	return &data
}
