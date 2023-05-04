package entity

type BetStatus string

const (
	PENDING BetStatus = "PENDING"
	SOLVED            = "SOLVED"
	LOST              = "LOST"
	CASHOUT           = "CASHOUT"
	AVOIDED           = "AVOIDED"
)

type CurrencyStatus string

const (
	DEPOSIT  CurrencyStatus = "DEPOSIT"
	WITHDRAW                = "WITHDRAW"
)
