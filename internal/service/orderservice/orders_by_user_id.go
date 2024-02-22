package orderservice

import (
	"context"
	"gophermart/internal/entity"
)

func (s *Service) GetByUserID(ctx context.Context, userID int64) ([]entity.Order, error) {
	return s.orderRepo.GetByUserID(ctx, userID)
}
