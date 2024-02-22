package orderrepo

import (
	"context"
	"gophermart/internal/entity"
)

func (r *Repository) UpdateStatusAndAccrual(ctx context.Context, orderID string, status entity.OrderStatus, accrual float64) error {
	_, err := r.Transactor.GetConn(ctx).Exec(
		"UPDATE orders SET status = $1, accrual = $2 WHERE order_id = $3",
		status, accrual, orderID,
	)

	return err
}
