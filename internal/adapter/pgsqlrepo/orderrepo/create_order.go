package orderrepo

import (
	"context"
	"gophermart/internal/entity"
	"time"
)

func (r *Repository) CreateOrder(ctx context.Context, orderID string, userID int64, status entity.OrderStatus, accrual float64) (int64, error) {
	var lastInsertID int64

	err := r.Transactor.GetConn(ctx).QueryRow(
		"INSERT INTO orders (order_id, user_id, status, accrual, created_at) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (order_id) DO NOTHING RETURNING id",
		orderID, userID, status, accrual, time.Now(),
	).Scan(&lastInsertID)

	if lastInsertID == 0 {
		return lastInsertID, &entity.DuplicateError{
			Field: "orderID",
			Value: orderID,
		}
	}

	if err != nil {
		return lastInsertID, err
	}

	return lastInsertID, nil
}
