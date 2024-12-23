package videoservice

import (
	"app/internal/connection"
	queuepayload "app/internal/dto/queue_payload"

	"gorm.io/gorm"
)

type videoService struct {
	psql *gorm.DB
}

type VideoService interface {
	UploadQuantityVideo(payload queuepayload.QueueUrlQuantityPayload) error
}

func Register() VideoService {
	return &videoService{
		psql: connection.GetPsql(),
	}
}
