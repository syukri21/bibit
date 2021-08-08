package http

import (
	"github.com/labstack/echo/v4"
)

func (r *HTTP) MovieRoute(g *echo.Group) {
	{
		path := g.Group("/movies")
		path.GET("", r.Controller.MovieHttp.Search)
	}
	{
		path := g.Group("/movie")
		path.GET("/:id", r.Controller.MovieHttp.GetDetail)
	}
}
