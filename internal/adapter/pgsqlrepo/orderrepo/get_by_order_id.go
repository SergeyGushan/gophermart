package orderrepo

import (
	"context"
	"gophermart/internal/entity"
)

func (r *Repository) GetByOrderID(ctx context.Context, orderID string) (entity.Order, error) {
	order := entity.Order{}

	err := r.Transactor.GetConn(ctx).QueryRow(
		"SELECT * FROM orders WHERE order_id = $1",
		orderID,
	).Scan(&order.ID, &order.OrderID, &order.UserID, &order.Status, &order.Accrual, &order.CreatedAt)

	return order, err
}
