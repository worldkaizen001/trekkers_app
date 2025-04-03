package models

import "gorm.io/gorm"

type Step struct {
	gorm.Model
	UserID uint
	step int
}
