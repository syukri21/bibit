package main

import (
	"github.com/labstack/echo/v4"
	"os"
)

func main() {

	_, found := os.LookupEnv("GRPC")

	if !found {
		e := echo.New()
		app := Module{}
		app.New(e).http.Run(e)
	} else {
		e := echo.New()
		app := Module{}
		app.New(e).grpc.Run()
	}

}
