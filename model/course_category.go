package model

import "gorm.io/gorm"

type CourseCategory struct {
	gorm.Model
	CourseId   uint
	CategoryId uint

	Course   *Course   `json:"course" gorm:"foreignKey:CourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category *Category `json:"category" gorm:"foreignKey:CategoryId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
