package controller

import (
	"net/http"
)

type authController struct {
}

type AuthControlle interface {
	Login(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request)        {}
func (c *authController) RefreshToken(w http.ResponseWriter, r *http.Request) {}
func (c *authController) Register(w http.ResponseWriter, r *http.Request)     {}

func NewAuthController() AuthControlle {
	return &authController{}
}
