package models

import "github.com/jinzhu/gorm"

type Car struct {
	gorm.Model
	Mfr              string
	Year             int
	Plate            int
	Gov              string
	UserID           int
	ImageFileID      int
	PlateImageFileID int
	StatusID         uint
}
