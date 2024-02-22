package entity

import (
	"encoding/json"
	"time"
)

type Order struct {
	ID        int64       `db:"id" json:"-"`
	OrderID   string      `db:"order_id" json:"number"`
	UserID    int64       `db:"user_id" json:"-"`
	Status    OrderStatus `db:"status" json:"status"`
	Accrual   float64     `db:"accrual" json:"accrual"`
	CreatedAt time.Time   `db:"created_at" json:"uploaded_at"`
}

func (o *Order) MarshalJSON() ([]byte, error) {
	type Alias Order
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"uploaded_at"`
	}{
		Alias:     (*Alias)(o),
		CreatedAt: o.CreatedAt.Format(time.RFC3339),
	})
}
