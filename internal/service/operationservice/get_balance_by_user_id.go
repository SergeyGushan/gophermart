package operationservice

import (
	"context"
	"gophermart/internal/entity"
)

func (s *Service) GetBalanceByUserID(ctx context.Context, userID int64) (entity.BalanceResponse, error) {
	return s.operationRepo.GetBalanceByUserID(ctx, userID)
}
