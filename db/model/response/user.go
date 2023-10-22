package response

type AdminLoginResponse struct {
	Id        uint   `json:"id"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

type UserLoginResponse struct {
	Id        uint   `json:"id"`
	HashCode  string `json:"hashcode"`
	Token     string `json:"token"`
	GameId    uint   `json:"game_id"`
	ExpiresAt int64  `json:"expiresAt"`
}
