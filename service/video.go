package service

import (
	"app/config"
	"app/constant"
	queuepayload "app/dto/queue_payload"
	"app/model"

	"gorm.io/gorm"
)

type videoService struct {
	psql *gorm.DB
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

	return nil
}

func NewVideoService() VideoService {
	return &videoService{
		psql: config.GetPsql(),
	}
}
