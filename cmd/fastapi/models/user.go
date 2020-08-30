package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name        string
	Mobile      string
	Email       string
	Password    string
	Gender      string
	ImageFileID uint
	IsDriver    bool
	StatusID    uint
}
