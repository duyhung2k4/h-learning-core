package videoservice

import (
	constant "app/internal/constants"
	queuepayload "app/internal/dto/queue_payload"
	"app/internal/entity"
)

func (s *videoService) UploadQuantityVideo(payload queuepayload.QueueUrlQuantityPayload) error {
	var newVideoLession entity.VideoLession

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

	err := s.psql.Model(&entity.VideoLession{}).Where("code = ?", payload.Uuid).Updates(&newVideoLession).Error
	if err != nil {
		return err
	}

	return nil
}
