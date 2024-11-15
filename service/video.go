package service

import (
	"app/config"
	"app/constant"
	queuepayload "app/dto/queue_payload"
	"app/model"
	"context"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type videoService struct {
	psql         *gorm.DB
	connRabbitmq *amqp091.Connection
}

type VideoService interface {
	UploadQuantityVideo(payload queuepayload.QueueUrlQuantityPayload) error
}

func (s *videoService) UploadQuantityVideo(payload queuepayload.QueueUrlQuantityPayload) error {
	var newVideoLession model.VideoLession

	switch payload.Quantity {
	case string(constant.QUANTITY_VIDEO_360P):
		newVideoLession.Url360p = &payload.Url
	case string(constant.QUANTITY_VIDEO_480P):
		newVideoLession.Url480p = &payload.Url
	case string(constant.QUANTITY_VIDEO_720P):
		newVideoLession.Url720p = &payload.Url
	case string(constant.QUANTITY_VIDEO_1080P):
		newVideoLession.Url1080p = &payload.Url
	}

	err := s.psql.Model(&model.VideoLession{}).Where("code = ?", payload.Uuid).Updates(&newVideoLession).Error
	if err != nil {
		return err
	}

	var videoLession model.VideoLession
	err = s.psql.Model(&model.VideoLession{}).Where("code = ?", payload.Uuid).First(&videoLession).Error
	if err != nil {
		return err
	}

	if videoLession.Url360p == nil {
		return nil
	}

	// if videoLession.Url360p == nil ||
	// 	videoLession.Url480p == nil ||
	// 	videoLession.Url720p == nil ||
	// 	videoLession.Url1080p == nil {
	// 	return nil
	// }

	ch, err := s.connRabbitmq.Channel()
	if err != nil {
		return err
	}

	payloadMess := queuepayload.QueueFileDeleteMp4{
		Uuid: payload.Uuid,
	}

	payloadJsonString, err := json.Marshal(payloadMess)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(context.Background(),
		"",
		string(constant.QUEUE_DELETE_MP4),
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        payloadJsonString,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func NewVideoService() VideoService {
	return &videoService{
		psql:         config.GetPsql(),
		connRabbitmq: config.GetRabbitmq(),
	}
}
