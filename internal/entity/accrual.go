package entity

type Accrual struct {
	Code int
	Item struct {
		Order   string  `json:"order"`
		Status  string  `json:"status"`
		Accrual float64 `json:"accrual"`
	}
}
