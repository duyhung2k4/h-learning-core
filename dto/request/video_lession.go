package request

type CreateVideoLessionReq struct {
	LessionId uint `json:"lessionId"`
}

type CheckVideoUploadReq struct {
	VideoLessionId uint `json:"videoLessionId"`
}
