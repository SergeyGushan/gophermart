package usecase

import (
	"context"
	"gophermart/internal/entity"
)

func (u *UseCase) GetOrdersByUserID(ctx context.Context, userID int64) ([]entity.Order, error) {
	return u.orderService.GetByUserID(ctx, userID)
}
