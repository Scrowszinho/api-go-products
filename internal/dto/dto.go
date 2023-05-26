package dto

import "time"

type CreateUserInput struct {
	Name     string  `json:"name"`
	LastName string  `json:"last_name"`
	Email    string  `json:"email"`
	Nickname string  `json:"nickname"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

type GetJWTInput struct {
	EmailOrNickname string `json:"user"`
	Password        string `json:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}

type CreateEventInput struct {
	Name        string    `json:"name"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
}
