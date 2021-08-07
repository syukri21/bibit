package main

import (
	"bibit-test/src/constants"
	"bibit-test/src/utils/proxy"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"os"
)

type Module struct{}

func (m *Module) New(e *echo.Echo) {
	ioc := m.NewIOC()
	s := NewAppService(ioc)
	s.NewRoute(e)
}

func (m *Module) NewIOC() di.Container {
	builder, _ := di.NewBuilder()

	_ = builder.Add(di.Def{
		Name: constants.PROXY_OMDBAPI,
		Build: func(ctn di.Container) (interface{}, error) {
			url, _ := os.LookupEnv("PROXY_OMDBAPI_URL")
			credential, _ := os.LookupEnv("PROXY_OMDBAPI_CREDENTIAL")
			return proxy.Create(url, credential), nil
		},
	})

	return builder.Build()
}
