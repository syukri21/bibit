package controllers

import (
	"bibit-test/src/models"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"net/http"
)

type MovieHTTP interface {
	Search(ctx echo.Context) error
}

type MovieHTTPImpl struct {
}

func (m MovieHTTPImpl) Search(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.GenericRes{
		Code:    http.StatusOK,
		Message: "Good",
		Data:    nil,
	})
}

func NewMovieHTTP(ioc di.Container) MovieHTTP {
	return &MovieHTTPImpl{}
}
