package models

import "gorm.io/gorm"

type Module struct {
	gorm.Model
	Name     string
	Semester int
	UserId   int
}
