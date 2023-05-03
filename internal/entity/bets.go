package entity

type Bets struct {
	ID          int     `gorm:"primaryKey"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	Status      Status  `json:"status"`
	Odd         float64 `json:"dds"`
	FinalOdd    float64 `json:"final_odd"`
}

func NewBet(name string, typee string, description string, status Status, odd float64, finalOdd float64, val float64) (*Bets, error) {
	bet := &Bets{
		Name:        name,
		Type:        typee,
		Description: description,
		Value:       val,
		Status:      status,
		Odd:         odd,
		FinalOdd:    finalOdd,
	}
	return bet, nil
}
