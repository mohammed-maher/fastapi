package models

import "github.com/jinzhu/gorm"

type File struct {
	gorm.Model
	Bucket   string
	Object   string
	Mime     string
	FileName string
	Ext      string
}
