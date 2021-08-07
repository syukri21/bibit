package controllers

import (
	"bibit-test/src/constants"
	"bibit-test/src/models/model"
	"bibit-test/src/models/proto"
	"bibit-test/src/services"
	"context"
	"encoding/json"
	"github.com/sarulabs/di"
	"net/http"
)

type MovieGRPCImpl struct {
	service *services.Service
}

func (m MovieGRPCImpl) Search(_ context.Context, params *proto.MoviesSearchParams) (payloads *proto.MoviesSearchResult, err error) {
	result, _, err := m.service.Movie.Search(model.MoviesSearchParams{
		Search: params.Search,
		Page:   int(params.Page),
	})

	if err != nil {
		return nil, err
	}

	data := &[]*proto.Search{}

	jsonByte, err := json.Marshal(result.Search)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonByte, data)
	if err != nil {
		return
	}

	return &proto.MoviesSearchResult{
		Data:    *data,
		Code:    int64(http.StatusOK),
		Message: "OK",
		Meta: &proto.MetaMoviesSearchResult{
			TotalResults: result.TotalResults,
			Response:     bool(result.Response),
		},
	}, nil
}

func (m MovieGRPCImpl) GetDetail(_ context.Context, params *proto.MoviesGetDetailParams) (payload *proto.MovieDetailResult, err error) {
	result, _, err := m.service.Movie.GetDetail(model.MoviesGetDetailParams{
		ID: params.ID,
	})
	if err != nil {
		return nil, err
	}

	data := &proto.MovieDetailResponse{}

	jsonByte, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonByte, data)
	if err != nil {
		return nil, err
	}
	return &proto.MovieDetailResult{
		Data:    data,
		Code:    int64(http.StatusOK),
		Message: "OK",
		Meta:    nil,
	}, nil

}

func NewMovieGRPC(ioc di.Container) proto.MoviesServer {
	return &MovieGRPCImpl{
		service: ioc.Get(constants.SERVICE).(*services.Service),
	}
}
