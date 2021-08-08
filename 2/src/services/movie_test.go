package services

import (
	"bibit-test/src/constants"
	"bibit-test/src/models/model"
	"bibit-test/src/repositories"
	"bibit-test/src/utils/fixtures"
	"errors"
	"github.com/sarulabs/di"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMovieSearchService(t *testing.T) {
	assertions := assert.New(t)
	t.Run("It should get data from ioc", func(t *testing.T) {
		builder, _ := di.NewBuilder()
		_ = builder.Add(di.Def{
			Name: constants.REPOSITORY,
			Build: func(ctn di.Container) (interface{}, error) {
				return repositories.NewRepositoryMock(), nil
			},
		})

		accountService := NewMovie(builder.Build())
		assertions.NotNil(t, accountService)
	})
}

func TestMovieSearch(t *testing.T) {
	t.Run("It should search movie successfully", func(t *testing.T) {
		mocks := repositories.MovieMock{
			SearchData:    &fixtures.MovieSearchResponse,
			GetDetailData: &fixtures.MoviesGetDetailResponse,
		}

		movieLogMock := repositories.MovieLogMock{
			LogCreateError: errors.New("Failed"),
		}

		repository := &repositories.Repository{
			Movie:    repositories.NewMovieMock(mocks),
			MovieLog: repositories.NewMovieLogMock(movieLogMock),
		}

		service := &MovieImpl{
			repository: repository,
		}

		data, _, _ := service.Search(model.MoviesSearchParams{
			Search: "Hello",
			Page:   1,
		})

		assert.Equal(t, data, &fixtures.MovieSearchResponse)
	})
	t.Run("It should get detail successfully", func(t *testing.T) {
		mocks := repositories.MovieMock{
			SearchData:    &fixtures.MovieSearchResponse,
			GetDetailData: &fixtures.MoviesGetDetailResponse,
		}

		movieLogMock := repositories.MovieLogMock{
			LogCreateError: errors.New("Failed"),
		}

		repository := &repositories.Repository{
			Movie:    repositories.NewMovieMock(mocks),
			MovieLog: repositories.NewMovieLogMock(movieLogMock),
		}

		service := &MovieImpl{
			repository: repository,
		}

		data, _, _ := service.GetDetail(model.MoviesGetDetailParams{
			ID: "t11188",
		})
		assert.Equal(t, data, &fixtures.MoviesGetDetailResponse)

	})
}
