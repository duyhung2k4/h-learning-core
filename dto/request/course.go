package request

type CreateCourseReq struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MultiLogin  bool    `json:"multiLogin"`
	Value       float64 `json:"value"`
}

type UpdateCourseReq struct {
	Id          uint     `json:"id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	MultiLogin  *bool    `json:"multiLogin"`
	Value       *float64 `json:"value"`
}

type ChangeAvticeCourseReq struct {
	Id     uint `json:"id"`
	Active bool `json:"active"`
}
