package entity

import (
	"encoding/json"
	"time"
)

type Operation struct {
	ID        int64         `db:"id" json:"-"`
	OrderID   string        `db:"order_id" json:"order"`
	UserID    int64         `db:"user_id" json:"-"`
	Sum       float64       `db:"sum" json:"sum"`
	Type      OperationType `db:"type" json:"-"`
	CreatedAt time.Time     `db:"created_at"  json:"uploaded_at"`
}

func (o *Operation) MarshalJSON() ([]byte, error) {
	type Alias Operation
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"uploaded_at"`
	}{
		Alias:     (*Alias)(o),
		CreatedAt: o.CreatedAt.Format(time.RFC3339),
	})
}
