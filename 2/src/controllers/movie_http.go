package controllers

import (
	"bibit-test/src/constants"
	"bibit-test/src/models"
	"bibit-test/src/services"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"net/http"
	"strconv"
)

type MovieHTTP interface {
	Search(ctx echo.Context) error
	GetDetail(ctx echo.Context) error
}

type MovieHTTPImpl struct {
	service *services.Service
}

func (m MovieHTTPImpl) GetDetail(ctx echo.Context) error {
	id := ctx.Param("id")
	result, code, err := m.service.Movie.GetDetail(models.MoviesGetDetailParams{
		ID: id,
	})

	if err != nil {
		return echo.NewHTTPError(code, err.Error())
	}

	return ctx.JSON(http.StatusOK, models.GenericRes{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    result,
		Meta:    nil,
	})

}

func (m MovieHTTPImpl) Search(ctx echo.Context) error {

	search := ctx.QueryParam("search")
	page := ctx.QueryParam("page")

	pageInt, err := strconv.ParseInt(page, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Parameters Page Invalid!!!")
	}

	result, code, err := m.service.Movie.Search(models.MoviesSearchParams{
		Search: search,
		Page:   int(pageInt),
	})

	if err != nil {
		return echo.NewHTTPError(code, err.Error())
	}

	return ctx.JSON(http.StatusOK, models.GenericRes{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    result.Search,
		Meta: map[string]interface{}{
			"response":    result.Response,
			"totalResult": result.TotalResults,
		},
	})
}

func NewMovieHTTP(ioc di.Container) MovieHTTP {
	return &MovieHTTPImpl{
		service: ioc.Get(constants.SERVICE).(*services.Service),
	}
}
