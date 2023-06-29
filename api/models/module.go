package models

import "gorm.io/gorm"

type Module struct {
	gorm.Model
	Name     string `json:"name"`
	Semester int `json:"semester"`
	UserId   uint `json:"userId"`
}

type ModuleDTO struct {
	Id       uint `json:"id"`
	Name     string `json:"name"`
	Semester int `json:"semester"`
	Grades   []GradeDTO `json:"grades"`
}
