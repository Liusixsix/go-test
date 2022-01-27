package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string `grom:"type:varchar(20);not null"`
	Telephone string `grom:"type:varchar(20);not null"`
	Password  string `grom:"size:255;not null"`
}
