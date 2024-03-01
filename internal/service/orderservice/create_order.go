package orderservice

import (
	"context"
	"gophermart/internal/entity"
)

func (s *Service) CreateOrder(ctx context.Context, orderID string, userID int64) (int64, error) {
	return s.orderRepo.CreateOrder(ctx, orderID, userID, entity.OrderStatusNew, 0)
}
