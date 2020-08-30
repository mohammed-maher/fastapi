package models

import "github.com/jinzhu/gorm"

type City struct {
	gorm.Model
	Name     string
	Long     float64
	Lat      float64
	StatusID uint
}
