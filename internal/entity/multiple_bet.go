package entity

type betsMul struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	Status      Status  `json:"status"`
	Odd         float64 `json:"dds"`
	SingleBet   bool    `json:"single_bet"`
}

type MultipleBet struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       float64
	Status      Status `json:"status"`
	Bets        []betsMul
	Odd         float64 `json:"odds"`
	FinalOdd    float64 `json:"final_odd"`
	Bonus       float64 `json:"bonus"`
}

func CreateMultipleBet(name string, description string, value float64, status Status, bets []betsMul, odd float64, finalOdd float64, bonus float64) (*MultipleBet, error) {
	mulbet := &MultipleBet{
		Name:        name,
		Description: description,
		Value:       value,
		Status:      status,
		Bets:        bets,
		Odd:         odd,
		FinalOdd:    finalOdd,
		Bonus:       bonus,
	}
	return mulbet, nil
}
