package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	Name     string `json:"name"`
	Grade    float32 `json:"grade"`
	ModuleId uint `json:"moduleId"`
	UserId   uint `json:"userId"`
}

type GradeDTO struct {
	Id       uint `json:"id"`
	Name     string `json:"name"`
	Grade    float32 `json:"grade"`
	ModuleId uint 	`json:"moduleId"`
}
