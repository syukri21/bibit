package http

import (
	"github.com/labstack/echo/v4"
)

func (r *HTTP) MovieRoute(g *echo.Group) {
	{
		path := g.Group("/movies")
		path.GET("", r.controller.MovieHttp.Search)
	}
	{
		path := g.Group("/movie")
		path.GET("/:id", r.controller.MovieHttp.GetDetail)
	}
}
