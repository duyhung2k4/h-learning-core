package model

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       int    `json:"order" gorm:"index:idx_order_chapter"`

	CourseId uint `json:"courseId" gorm:"index:idx_order_chapter"`

	Course   *Course   `json:"course" gorm:"foreignKey:CourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Lessions []Lession `json:"lessions" gorm:"foreignKey:ChapterId"`
}
