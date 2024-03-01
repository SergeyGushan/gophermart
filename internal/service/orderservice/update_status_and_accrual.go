package orderservice

import (
	"context"
	"gophermart/internal/entity"
)

func (s *Service) UpdateStatusAndAccrual(ctx context.Context, orderID string, status entity.OrderStatus, accrual float64) error {
	return s.orderRepo.UpdateStatusAndAccrual(ctx, orderID, status, accrual)
}
