package usecase

import (
	"context"
	"gophermart/internal/entity"
)

func (u *UseCase) GetBalanceByUserID(ctx context.Context, userID int64) (entity.BalanceResponse, error) {
	return u.operationService.GetBalanceByUserID(ctx, userID)
}
