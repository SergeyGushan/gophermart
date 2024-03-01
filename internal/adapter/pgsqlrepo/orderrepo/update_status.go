package orderrepo

import (
	"context"
	"gophermart/internal/entity"
)

func (r *Repository) UpdateStatus(ctx context.Context, orderID string, status entity.OrderStatus) error {
	_, err := r.Transactor.GetConn(ctx).Exec(
		"UPDATE orders SET status = $1 WHERE order_id = $2",
		status, orderID,
	)

	return err
}
