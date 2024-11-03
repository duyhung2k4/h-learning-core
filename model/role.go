package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name"`
	Code string `json:"code" gorm:"unique"`

	Profiles []Profile `json:"profiles" gorm:"foreignKey:RoleId;"`
}
