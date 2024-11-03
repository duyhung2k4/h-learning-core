package model

import "gorm.io/gorm"

type VideoLession struct {
	gorm.Model
	Code     string `json:"code" gorm:"unique"`
	Url      string `json:"url"`
	Thumnail string `json:"thumnail"`

	LessionId uint     `json:"lessionId"`
	Lession   *Lession `json:"lession" gorm:"foreignKey:LessionId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
