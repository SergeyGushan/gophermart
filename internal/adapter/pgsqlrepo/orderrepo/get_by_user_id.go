package orderrepo

import (
	"context"
	"database/sql"
	"gophermart/internal/entity"
	"log"
)

func (r *Repository) GetByUserID(ctx context.Context, userID int64) ([]entity.Order, error) {
	orders := make([]entity.Order, 0)

	rows, err := r.Transactor.GetConn(ctx).Query(
		"SELECT * FROM orders WHERE user_id = $1",
		userID,
	)

	if err != nil {
		return orders, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	if rows.Err() != nil {
		return orders, rows.Err()
	}

	for rows.Next() {
		order := entity.Order{}

		err := rows.Scan(&order.ID, &order.OrderID, &order.UserID, &order.Status, &order.Accrual, &order.CreatedAt)
		if err != nil {
			continue
		}

		orders = append(orders, order)
	}

	return orders, nil
}
