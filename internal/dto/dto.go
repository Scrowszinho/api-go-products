package dto

import (
	"time"

	"github.com/Scrowszinho/api-go-products/internal/entity"
)

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
	*entity.User   `json:"user"`
	AccessToken    string    `json:"access_token"`
	CreationDate   time.Time `json:"creation_date"`
	ExpirationDate time.Time `json:"expiration_date"`
}

type CreateEventInput struct {
	Name        string    `json:"name"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
}

type CreateOutcomesInput struct {
	Name    string           `json:"name"`
	Odds    float64          `json:"odds"`
	EventID int              `json:"event_id"`
	Status  entity.BetStatus `json:"status"`
}

type CreateBetsInput struct {
	User      entity.User `json:"user"`
	OutcomeID int         `json:"outcome_id"`
	Amount    float64     `json:"amount"`
	Active    bool        `json:"active"`
	Bonus     float64     `json:"bonus"`
}
