package entity

type Status string

const (
	PENDING  Status = "PENDING"
	SOLVED          = "SOLVED"
	LOST            = "LOST"
	PARTIAL         = "PARTIAL"
	TURNBACK        = "TURNBACK"
	SOLDED          = "SOLDED"
)
