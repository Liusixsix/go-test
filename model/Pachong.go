package model

import "github.com/jinzhu/gorm"

type Pachong struct {
	gorm.Model
	Name string `grom:"type:varchar(255);not null"`
	Href string `grom:"type:varchar(255);"`
}
