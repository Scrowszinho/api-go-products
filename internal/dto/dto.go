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
	Email    string `json:"email"`
	Password string `json:"password"`
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
	Name     string           `json:"name"`
	Odds     float64          `json:"odds"`
	EventID  int              `json:"event_id"`
	Status   entity.BetStatus `json:"status"`
	MarketID int              `json:"market_id"`
}

type CreateBetsInput struct {
	OutcomeID int     `json:"outcome_id"`
	Amount    float64 `json:"amount"`
	Active    bool    `json:"active"`
	Bonus     float64 `json:"bonus"`
	Odds      float64 `json:"odds"`
}

type GetSingleBetsByUser struct {
	ID               uint      `json:"id"`
	Amount           float64   `json:"amount"`
	Odds             float64   `json:"odds"`
	Status           string    `json:"status"`
	OutcomeName      string    `json:"name"`
	EventID          uint      `json:"event_id"`
	EventName        string    `json:"event_name"`
	EventStartTime   time.Time `json:"start_time"`
	EventEndTime     time.Time `json:"end_time"`
	EventDescription string    `json:"event_description"`
}
