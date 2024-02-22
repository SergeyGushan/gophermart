package operationservice

import (
	"context"
	"gophermart/internal/entity"
)

func (s *Service) CreateOperation(ctx context.Context, orderID string, userID int64, operationType entity.OperationType, sum float64) (int64, error) {
	return s.operationRepo.CreateOperation(ctx, orderID, userID, operationType, sum)
}
