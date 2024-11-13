package response

import (
	"app/model"
	"time"
)

type RegisterRes struct {
	Token string    `json:"token"`
	Exp   time.Time `json:"exp"`
}

type LoginRes struct {
	AccessToken  string        `json:"accessToken"`
	RefreshToken string        `json:"refreshToken"`
	Profile      model.Profile `json:"profile"`
}

type RefreshTokenRes struct {
	AccessToken  string        `json:"accessToken"`
	RefreshToken string        `json:"refreshToken"`
	Profile      model.Profile `json:"profile"`
}
