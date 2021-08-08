package repositories

import (
	"bibit-test/src/models/model"
	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type MovieLog interface {
	CreateSearchLog(data *model.MovieSearchLog) (err error)
	CreateGetDetailLog(data *model.MovieGetDetailLog) (err error)
}

type MovieLogImpl struct {
	db *gorm.DB
}

func (m MovieLogImpl) NewDb() *gorm.DB {
	return m.db.Session(&gorm.Session{})
}

func (m MovieLogImpl) CreateGetDetailLog(data *model.MovieGetDetailLog) (err error) {
	return m.NewDb().Model(&data).Create(&data).Error
}

func (m MovieLogImpl) CreateSearchLog(data *model.MovieSearchLog) (err error) {
	return m.NewDb().Model(&data).Create(&data).Error
}

func NewMovieLog(ioc di.Container) MovieLog {
	return &MovieLogImpl{
		db: getDatabase(ioc),
	}
}
