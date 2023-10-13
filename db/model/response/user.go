package response

type LoginResponse struct {
	Id        uint   `json:"id"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}
