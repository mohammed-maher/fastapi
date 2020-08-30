package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	TripID     uint
	Passengers uint
	Confirmed  bool
	UserID     uint
	TwoWay     bool
	StatusID   uint
}
