package orderservice

import (
	"context"
	"gophermart/internal/entity"
)

func (s *Service) UpdateStatus(ctx context.Context, orderID string, status entity.OrderStatus) error {
	return s.orderRepo.UpdateStatus(ctx, orderID, status)
}
