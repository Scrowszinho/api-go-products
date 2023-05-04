package entity

type Status string

const (
	PENDING Status = "PENDING"
	SOLVED         = "SOLVED"
	LOST           = "LOST"
	CASHOUT        = "CASHOUT"
	AVOIDED        = "AVOIDED"
)
