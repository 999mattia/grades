package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	Name     string
	Grade    float32
	ModuleId uint
	UserId   uint
}

type GradeDTO struct {
	Id       uint
	Name     string
	Grade    float32
	ModuleId uint
}
