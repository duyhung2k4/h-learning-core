package config

import (
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/time/rate"
)

func makeVariable() {
	mapSocket = make(map[string]*websocket.Conn)
	mapSocketEvent = make(map[string]map[string]*websocket.Conn)

	// chanel job
	emailChan = make(chan EmailJob_MessPayload)

	// http
	limiter = rate.NewLimiter(rate.Every(time.Second), 500)
}
