package request

type CreateChapterReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CourseId    uint   `json:"courseId"`
}

type UpdateChapterReq struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteChapterReq struct {
	Id uint `json:"id"`
}
