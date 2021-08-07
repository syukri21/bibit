package http

import (
	"github.com/labstack/echo/v4"
)

func (r *Route) MovieRoute(g *echo.Group) {
	{
		g.GET("/movies", r.controller.MovieHttp.Search)
	}
}
