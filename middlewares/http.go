package middlewares

import (
	"app/controller"
	"context"
	"net/http"

	"github.com/go-chi/render"
	"golang.org/x/time/rate"
)

func authServerError(w http.ResponseWriter, r *http.Request, err error) {
	res := controller.Response{
		Data:    nil,
		Message: err.Error(),
		Status:  401,
		Error:   err,
	}
	w.WriteHeader(http.StatusUnauthorized)
	render.JSON(w, r, res)
}

func RateLimiter(limiter *rate.Limiter) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// if limiter.Allow() == false {
			// 	http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			// 	return
			// }
			if err := limiter.Wait(context.Background()); err != nil {
				http.Error(w, "Error occurred while waiting for rate limit", http.StatusInternalServerError)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
