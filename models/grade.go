package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	Name     string
	Grade    float32
	ModuleId int
	UserId   int
}
