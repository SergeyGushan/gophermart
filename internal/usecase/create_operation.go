package usecase

import (
	"context"
	"gophermart/internal/entity"
)

func (u *UseCase) CreateOperation(ctx context.Context, orderID string, userID int64, operationType entity.OperationType, sum float64) (int64, error) {
	return u.operationService.CreateOperation(ctx, orderID, userID, operationType, sum)
}
