package entity

import (
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
	Code string
}
