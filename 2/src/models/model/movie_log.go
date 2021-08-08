package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type MovieSearchLog struct {
	gorm.Model
	ID   uuid.UUID `gorm:"column:id" json:"id"`
	Data JSON      `gorm:"column:data;type:json" json:"data"`
}

func (e *MovieSearchLog) BeforeCreate(*gorm.DB) (err error) {
	e.ID = uuid.NewV4()
	return
}

type MovieGetDetailLog struct {
	gorm.Model
	ID   uuid.UUID `gorm:"column:id" json:"id"`
	Data JSON      `gorm:"column:data;type:json" json:"data"`
}

func (e *MovieGetDetailLog) BeforeCreate(*gorm.DB) (err error) {
	e.ID = uuid.NewV4()
	return
}
