package dto

type CreateUserInput struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	LastName string  `json:"last_name"`
	Email    string  `json:"email"`
	Nickname string  `json:"nickname"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

type GetJWTInput struct {
	EmailOrNickname string `json:"email"`
	Password        string `json:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}
