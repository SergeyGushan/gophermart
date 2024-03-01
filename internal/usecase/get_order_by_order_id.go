package usecase

import (
	"context"
	"gophermart/internal/entity"
)

func (u *UseCase) GetOrderByOrderID(ctx context.Context, orderID string) (entity.Order, error) {
	return u.orderService.GetByOrderID(ctx, orderID)
}
