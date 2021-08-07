package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	e := echo.New()

	app := Module{}

	app.New(e)

	port, found := os.LookupEnv("PORT")

	if !found {
		port = "1323"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
