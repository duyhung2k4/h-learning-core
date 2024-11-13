package request

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AcceptCode struct {
	Code string `json:"code"`
}
