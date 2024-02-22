package usecase

import (
	"context"
	"gophermart/internal/entity"
)

func (u *UseCase) GetOperationsByTypeWithdrawByUserID(ctx context.Context, userID int64) ([]entity.Operation, error) {
	return u.operationService.GetOperationsByTypeWithdrawByUserID(ctx, userID)
}
