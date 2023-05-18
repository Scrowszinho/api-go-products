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
	DEPOSIT     CurrencyStatus = "DEPOSIT"
	WITHDRAW    CurrencyStatus = "WITHDRAW"
	SUCCESS_BET CurrencyStatus = "SUCCESS_BET"
	LOST_BET    CurrencyStatus = "LOST_BET"
	AVOIDED_BET CurrencyStatus = "AVOIDED_BET"
)
