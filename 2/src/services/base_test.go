package services

import (
	"bibit-test/src/constants"
	"bibit-test/src/repositories"
	"github.com/sarulabs/di"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewService(t *testing.T) {
	assertions := assert.New(t)
	t.Run("It should passing ios into service", func(t *testing.T) {
		builder, _ := di.NewBuilder()
		_ = builder.Add(di.Def{
			Name: constants.REPOSITORY,
			Build: func(ctn di.Container) (interface{}, error) {
				return repositories.NewRepositoryMock(), nil
			},
		})

		service := NewService(builder.Build())
		assertions.NotNil(t, service)
	})
}
