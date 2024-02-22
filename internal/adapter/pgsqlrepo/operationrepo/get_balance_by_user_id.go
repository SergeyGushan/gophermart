package operationrepo

import (
	"context"
	"gophermart/internal/entity"
)

func (r *Repository) GetBalanceByUserID(ctx context.Context, userID int64) (entity.BalanceResponse, error) {
	var balance entity.BalanceResponse

	row := r.Transactor.GetConn(ctx).QueryRow(
		`SELECT 
			accrual - withdrawn as current,
			withdrawn
		FROM (
		    SELECT 
				COALESCE(SUM(CASE WHEN type = $1 THEN sum ELSE 0 END), 0) as withdrawn,
				COALESCE(SUM(CASE WHEN type = $2 THEN sum ELSE 0 END), 0) as accrual
			FROM operations 
			WHERE user_id = $3
		) as o`,
		entity.OperationTypeWithdrawn,
		entity.OperationTypeAccrual,
		userID,
	)

	if row.Err() != nil {
		return balance, row.Err()
	}

	err := row.Scan(&balance.Current, &balance.Withdrawn)

	if err != nil {
		return balance, err
	}

	return balance, nil
}
