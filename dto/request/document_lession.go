package request

type CreateDocumentLessionReq struct {
	Content   string `json:"content"`
	LessionId uint   `json:"lessionId"`
}

type UpdateDocumentLessionReq struct {
	Id      uint   `json:"id"`
	Content string `json:"content"`
}

type DeleteDocumentLessionReq struct {
	Id uint `json:"id"`
}
