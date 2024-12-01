package request

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type AcceptCode struct {
	Code string `json:"code"`
}
