package model

import "gorm.io/gorm"

type QuantityVideoLession struct {
	gorm.Model
	Url      string `json:"url"`
	Quantity string `json:"quantity"`

	VideoLessionId uint `json:"videoLessionId"`

	VideoLession *VideoLession `json:"videoLession" gorm:"foreignKey:VideoLessionId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
