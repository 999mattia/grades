package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique" json:"name"` 
	Password string `json:"password"`
}

type UserDTO struct {
	Id      uint `json:"id"` 
	Name    string `json:"name"`
	Modules []ModuleDTO `json:"modules"`
}
