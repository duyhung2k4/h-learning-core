package utils

import (
	"app/config"
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/jwtauth/v5"
)

type jwtUtils struct {
	jwt *jwtauth.JWTAuth
}

type JwtUtils interface {
	GetToken(r *http.Request) string
	GetMapToken(token string) (map[string]interface{}, error)
	JwtEncode(data map[string]interface{}) (string, error)
	JwtDecode(tokenString string) (map[string]interface{}, error)
}

func (j *jwtUtils) GetToken(r *http.Request) string {
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	return tokenString
}

func (j *jwtUtils) GetMapToken(token string) (map[string]interface{}, error) {
	mapData, errMapData := j.JwtDecode(token)
	if errMapData != nil {
		return nil, errMapData
	}

	return mapData, nil
}

func (j *jwtUtils) JwtEncode(data map[string]interface{}) (string, error) {
	_, tokenString, err := j.jwt.Encode(data)
	return tokenString, err
}

func (j *jwtUtils) JwtDecode(tokenString string) (map[string]interface{}, error) {
	var dataMap map[string]interface{}
	jwt, err := j.jwt.Decode(tokenString)
	if err != nil {
		return dataMap, err
	}

	dataMap, errMap := jwt.AsMap(context.Background())
	return dataMap, errMap
}

func NewJwtUtils() JwtUtils {
	return &jwtUtils{
		jwt: config.GetJWT(),
	}
}
