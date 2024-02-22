package orderservice

import (
	"context"
	"gophermart/internal/entity"
)

func (s *Service) GetByOrderID(ctx context.Context, orderID string) (entity.Order, error) {
	return s.orderRepo.GetByOrderID(ctx, orderID)
}
