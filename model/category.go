package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name"`
	Code string `json:"code"`

	CourseCategorys []CourseCategory `json:"courseCategorys" gorm:"foreignKey:CategoryId;"`
}
