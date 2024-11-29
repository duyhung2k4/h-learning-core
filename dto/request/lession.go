package request

type CreateLessionReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CourseId    uint   `json:"courseId"`
	ChapterId   uint   `json:"chapterId"`
}

type UpdateLessionReq struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteLessionReq struct {
	Id uint `json:"id"`
}
