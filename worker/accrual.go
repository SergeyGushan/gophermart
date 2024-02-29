package worker

import (
	"database/sql"
	"gophermart/internal/adapter/httprepo/accrualrepo"
	"gophermart/internal/entity"
	"log"
	"net/http"
)

func Accrual(pgSQL *sql.DB, repo *accrualrepo.Repository) {
	rows, err := pgSQL.Query(
		"SELECT * FROM orders WHERE status = $1 OR status = $2 LIMIT 5",
		entity.OrderStatusNew, entity.OrderStatusError,
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	if rows.Err() != nil {
		log.Fatal(rows.Err())
		return
	}

	for rows.Next() {
		order := entity.Order{}

		err := rows.Scan(&order.ID, &order.OrderID, &order.UserID, &order.Status, &order.Accrual, &order.CreatedAt)
		if err != nil {
			continue
		}

		_, err = pgSQL.Exec(
			"UPDATE orders SET status = $1 WHERE order_id = $2",
			entity.OrderStatusProcessing,
			order.OrderID,
		)

		if err != nil {
			log.Println("Error updating order status to Processing:", err)
			continue
		}

		accrual, errAccrual := repo.GetOrderAccrual(order.OrderID)

		if errAccrual != nil {
			log.Println("Error getting order accrual:", errAccrual)
			_, _ = pgSQL.Exec(
				"UPDATE orders SET status = $1 WHERE order_id = $2",
				entity.OrderStatusError,
				order.OrderID,
			)
			continue
		}

		if accrual.Code == http.StatusOK {
			switch accrual.Item.Status {
			case "INVALID":
				_, err := pgSQL.Exec(
					"UPDATE orders SET status = $1 WHERE order_id = $2",
					entity.OrderStatusInvalid,
					order.OrderID,
				)
				if err != nil {
					log.Println("Error updating order status to INVALID:", err)
					continue
				}
			case "PROCESSED":
				_, err := pgSQL.Exec(
					"UPDATE orders SET status = $1, accrual = $2 WHERE order_id = $3",
					entity.OrderStatusProcessed,
					accrual.Item.Accrual,
					order.OrderID,
				)

				if err != nil {
					log.Println("Error updating order status to PROCESSED:", err)
					continue
				}

				_, err = pgSQL.Exec(
					"INSERT INTO operations (order_id, user_id, type, sum) VALUES ($1, $2, $3, $4)",
					order.OrderID,
					order.UserID,
					entity.OperationTypeAccrual,
					accrual.Item.Accrual,
				)

				if err != nil {
					log.Println("Error inserting operation:", err)
					continue
				}
			default:
				log.Println("Unknown accrual status:", accrual.Item.Status)
			}
		}
	}
}
