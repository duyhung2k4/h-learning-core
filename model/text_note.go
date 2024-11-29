package model

import "gorm.io/gorm"

type TextNote struct {
	gorm.Model
	Time    string `json:"time"`
	Context string `json:"context"`

	VideoLessionId uint `json:"videoLessionId"`
	CreateId       uint `json:"createId"`

	VideoLession *VideoLession `json:"videoLession" gorm:"foreignKey:VideoLessionId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Creater      *Profile      `json:"creater" gorm:"foreignKey:CreateId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
