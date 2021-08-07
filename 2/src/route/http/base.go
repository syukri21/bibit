package http

import (
	"bibit-test/src/constants"
	"bibit-test/src/controllers"
	middlewares "bibit-test/src/middlwares"
	"bibit-test/src/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"net/http"
)

type Route struct {
	ioc        di.Container
	controller *controllers.Controller
}

func NewRoute(ioc di.Container) *Route {
	return &Route{
		controller: ioc.Get(constants.CONTROLLER).(*controllers.Controller),
	}
}

func (r Route) Run(e *echo.Echo) {
	e.Validator = middlewares.NewValidator(validator.New())
	e.HTTPErrorHandler = middlewares.ErrorHandler
	g := e.Group("")

	r.MovieRoute(g)
	r.HealthRoute(g)
	r.NotFoundRoute(g)
}

func (*Route) HealthRoute(g *echo.Group) {
	{
		g.GET("/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, models.GenericRes{
				Code:    http.StatusOK,
				Message: "OK.",
			})
		})
	}
}

func (*Route) NotFoundRoute(g *echo.Group) {
	{
		g.Any("*", func(c echo.Context) error {
			return c.JSON(http.StatusOK, models.GenericRes{
				Code:    http.StatusNotFound,
				Message: "Route not found.",
			})
		})
	}
}
