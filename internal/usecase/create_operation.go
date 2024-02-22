package usecase

import (
	"context"
	"gophermart/internal/entity"
	"gophermart/internal/handler/http/api"
	"net/http"
)

func (u *UseCase) CreateWithdrawnOperation(ctx context.Context, orderID string, sum float64) (int64, error) {
	userID, errContext := api.GetUserIDFromContext(ctx)

	if errContext != nil {
		return 0, entity.NewHTTPException(http.StatusUnauthorized, "")
	}

	balance, errGetBalanceByUserID := u.operationService.GetBalanceByUserID(ctx, userID)

	if errGetBalanceByUserID != nil {
		return 0, entity.NewHTTPException(http.StatusBadRequest, "")
	}

	if balance.Current < sum {
		return 0, entity.NewHTTPException(http.StatusPaymentRequired, "")
	}

	return u.operationService.CreateOperation(ctx, orderID, userID, entity.OperationTypeWithdrawn, sum)
}
