package http

import (
	"bibit-test/src/constants"
	"bibit-test/src/controllers"
	middlewares "bibit-test/src/middlwares"
	"bibit-test/src/models/model"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"net/http"
	"os"
)

type HTTP struct {
	ioc        di.Container
	controller *controllers.Controller
}

func New(ioc di.Container) *HTTP {
	return &HTTP{
		controller: ioc.Get(constants.CONTROLLER).(*controllers.Controller),
	}
}

func (r *HTTP) Run(e *echo.Echo) {
	e.Validator = middlewares.NewValidator(validator.New())
	e.HTTPErrorHandler = middlewares.ErrorHandler
	g := e.Group("")

	r.MovieRoute(g)
	r.HealthRoute(g)
	r.NotFoundRoute(g)

	port, found := os.LookupEnv("PORT")
	if !found {
		port = "1323"
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func (*HTTP) HealthRoute(g *echo.Group) {
	{
		g.GET("/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, model.GenericRes{
				Code:    http.StatusOK,
				Message: "OK.",
			})
		})
	}
}

func (*HTTP) NotFoundRoute(g *echo.Group) {
	{
		g.Any("*", func(c echo.Context) error {
			return c.JSON(http.StatusOK, model.GenericRes{
				Code:    http.StatusNotFound,
				Message: "Route not found.",
			})
		})
	}
}
