package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Trip struct {
	gorm.Model
	CarID         uint
	FromCityID    uint
	ToCityID      uint
	Passengers    uint
	TwoWay        bool
	DepartureDate time.Time
	ReturnDate    *time.Time
	StatusID      uint
}
