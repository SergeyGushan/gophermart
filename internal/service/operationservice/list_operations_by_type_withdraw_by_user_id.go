package operationservice

import (
	"context"
	"gophermart/internal/entity"
)

func (s *Service) GetOperationsByTypeWithdrawByUserID(ctx context.Context, userID int64) ([]entity.Operation, error) {
	return s.operationRepo.GetOperationsByUserID(ctx, userID, entity.OperationTypeWithdrawn)
}
