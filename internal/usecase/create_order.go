package usecase

import (
	"context"
)

func (u *UseCase) CreateOrder(ctx context.Context, orderID string, userID int64) (int64, error) {
	return u.orderService.CreateOrder(ctx, orderID, userID)
}
