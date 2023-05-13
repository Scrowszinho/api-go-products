package entity

type BetStatus string

const (
	PENDING BetStatus = "PENDING"
	SOLVED  BetStatus = "SOLVED"
	LOST    BetStatus = "LOST"
	CASHOUT BetStatus = "CASHOUT"
	AVOIDED BetStatus = "AVOIDED"
)

type CurrencyStatus string

const (
	DEPOSIT  CurrencyStatus = "DEPOSIT"
	WITHDRAW BetStatus      = "WITHDRAW"
)
