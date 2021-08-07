package main

import (
	middlewares "bibit-test/src/middlwares"
	"bibit-test/src/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"net/http"
)

type AppService struct {
}

func NewAppService(ioc di.Container) AppService {
	return AppService{}
}

func (s AppService) NewRoute(e *echo.Echo) {
	e.Validator = middlewares.NewValidator(validator.New())
	e.HTTPErrorHandler = middlewares.ErrorHandler

	g := e.Group("")
	s.HealthRoute(g)
	s.NotFoundRoute(g)
}

func (*AppService) HealthRoute(g *echo.Group) {
	{
		g.GET("/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, models.GenericRes{
				Code:    http.StatusOK,
				Message: "OK.",
			})
		})
	}
}

func (*AppService) NotFoundRoute(g *echo.Group) {
	{
		g.Any("*", func(c echo.Context) error {
			return c.JSON(http.StatusOK, models.GenericRes{
				Code:    http.StatusNotFound,
				Message: "Route not found.",
			})
		})
	}
}
