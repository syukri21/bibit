package main

import (
	"bibit-test/src/constants"
	"bibit-test/src/controllers"
	"bibit-test/src/repositories"
	"bibit-test/src/services"
	"bibit-test/src/utils/proxy"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"os"
)

type Module struct{}

func (m *Module) New(*echo.Echo) {
	_ = godotenv.Load(".env")
	ioc := m.NewIOC()

	s := NewApp(ioc)
	//s.http.Run(e)
	s.grpc.Run()
}

func (m *Module) NewIOC() di.Container {
	builder, _ := di.NewBuilder()

	_ = builder.Add(di.Def{
		Name: constants.PROXY_OMDBAPI,
		Build: func(ctn di.Container) (interface{}, error) {
			host := os.Getenv("PROXY_OMDBAPI_URL")
			credential := os.Getenv("PROXY_OMDBAPI_CREDENTIAL")
			return proxy.Create(host, credential), nil
		},
	})

	_ = builder.Add(di.Def{
		Name: constants.REPOSITORY,
		Build: func(ctn di.Container) (interface{}, error) {
			return repositories.NewRepository(builder.Build()), nil
		},
	})

	_ = builder.Add(di.Def{
		Name: constants.SERVICE,
		Build: func(ctn di.Container) (interface{}, error) {
			return services.NewService(builder.Build()), nil
		},
	})

	_ = builder.Add(di.Def{
		Name: constants.CONTROLLER,
		Build: func(ctn di.Container) (interface{}, error) {
			return controllers.NewController(builder.Build()), nil
		},
	})

	return builder.Build()
}
