package usecase

import (
	"context"
	"gophermart/internal/entity"
	"gophermart/internal/handler/http/api"
	"net/http"
)

func (u *UseCase) GetBalanceByUserID(ctx context.Context) (entity.BalanceResponse, error) {
	userID, errContext := api.GetUserIDFromContext(ctx)
	var balance entity.BalanceResponse

	if errContext != nil {
		return balance, entity.NewHTTPException(http.StatusUnauthorized, "")
	}

	balance, errGetBalanceByUserID := u.operationService.GetBalanceByUserID(ctx, userID)

	if errGetBalanceByUserID != nil {
		return balance, entity.NewHTTPException(http.StatusBadRequest, "")
	}

	return balance, nil
}
