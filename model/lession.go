package model

import "gorm.io/gorm"

type Lession struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       int    `json:"order" gorm:"idx:idx_order_lession"`

	ChapterId uint `json:"chapterId" gorm:"idx:idx_order_lession"`
	CourseId  uint `json:"courseId" gorm:"idx:idx_order_lession"`

	Chapter      *Chapter      `json:"chapter" gorm:"foreignKey:ChapterId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course       *Course       `json:"course" gorm:"foreignKey:CourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VideoLession *VideoLession `json:"videoLession" gorm:"foreignKey:LessionId"`
}
