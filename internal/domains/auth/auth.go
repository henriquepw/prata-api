package auth

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
