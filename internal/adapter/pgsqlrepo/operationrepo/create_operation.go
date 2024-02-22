package operationrepo

import (
	"context"
	"gophermart/internal/entity"
	"time"
)

func (r *Repository) CreateOperation(ctx context.Context, orderID string, userID int64, operationType entity.OperationType, sum float64) (int64, error) {
	var lastInsertID int64

	err := r.Transactor.GetConn(ctx).QueryRow(
		"INSERT INTO operations (order_id, user_id, sum, type, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		orderID, userID, sum, operationType, time.Now(),
	).Scan(&lastInsertID)

	return lastInsertID, err
}
