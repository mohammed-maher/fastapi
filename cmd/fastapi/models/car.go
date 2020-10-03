package models

import "github.com/jinzhu/gorm"

type Car struct {
	gorm.Model
	Mfr              string
	Year             uint
	Name             string
	Plate            string
	Gov              string
	UserID           uint
	ImageFileID      uint
	PlateImageFileID uint
	StatusID         uint
}
