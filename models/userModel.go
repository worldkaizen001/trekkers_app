package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Profile Profile `gorm:"foreignKey:id"`
	Steps []Step `gorm:"foreignKey:id"`

}


