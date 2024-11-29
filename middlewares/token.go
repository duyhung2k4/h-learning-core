package middlewares

import (
	"app/utils"
	"errors"
	"net/http"
	"strings"
	"time"
)

type middlewares struct {
	utils utils.JwtUtils
}

type Middlewares interface {
	ValidateExpAccessToken() func(http.Handler) http.Handler
}

func (m *middlewares) ValidateExpAccessToken() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		funcHttp := func(w http.ResponseWriter, r *http.Request) {
			if len(strings.Split(r.Header.Get("Authorization"), " ")) != 2 {
				authServerError(w, r, errors.New("token not found"))
				return
			}

			tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
			mapData, errMapData := m.utils.JwtDecode(tokenString)

			if errMapData != nil {
				authServerError(w, r, errMapData)
				return
			}

			exp := mapData["exp"].(time.Time)

			if time.Now().Unix() > exp.Unix() {
				authServerError(w, r, errors.New("token expired"))
				return
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(funcHttp)
	}
}

func NewMiddlewares() Middlewares {
	return &middlewares{
		utils: utils.NewJwtUtils(),
	}
}
