package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserId uint
	Bio string
	Hobbies []string
	
}