package operationrepo

import (
	"context"
	"database/sql"
	"gophermart/internal/entity"
	"log"
)

func (r *Repository) GetOperationsByUserID(ctx context.Context, userID int64, operationType entity.OperationType) ([]entity.Operation, error) {
	operations := make([]entity.Operation, 0)

	rows, err := r.Transactor.GetConn(ctx).Query(
		"SELECT * FROM operations WHERE user_id = $1 AND type = $2",
		userID, operationType,
	)

	if err != nil {
		return operations, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	if rows.Err() != nil {
		return operations, rows.Err()
	}

	for rows.Next() {
		operation := entity.Operation{}

		err := rows.Scan(&operation.ID, &operation.OrderID, &operation.UserID, &operation.Sum, &operation.Type, &operation.CreatedAt)
		if err != nil {
			continue
		}

		operations = append(operations, operation)
	}

	return operations, nil
}
