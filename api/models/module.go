package models

import "gorm.io/gorm"

type Module struct {
	gorm.Model
	Name     string
	Semester int
	UserId   uint
}

type ModuleDTO struct {
	Id       uint
	Name     string
	Semester int
	Grades   []GradeDTO
}
