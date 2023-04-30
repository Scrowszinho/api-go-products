package entity

type Bets struct {
	ID          int     `gorm:primaryKey`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	Status      string  `json:"status"`
	Odd         float64 `json:"dds"`
	FinalOdd    float64 `json:"final_odd"`
}

type Multiple struct {
	ID       int `gorm:primaryKey`
	Value    float64
	Status   string `json:"status"`
	Bets     []Bets
	Odd      float64 `json:"odds"`
	FinalOdd float64 `json:"final_odd"`
	Bonus    float64 `json:"bonus"`
}
