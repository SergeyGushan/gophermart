package usecase

import (
	"context"
	"errors"
	"gophermart/internal/entity"
	"gophermart/internal/handler/http/api"
	"gophermart/internal/service/luhn"
	"net/http"
)

func (u *UseCase) CreateOrder(ctx context.Context, orderID string) (int64, error) {
	var duplicateError *entity.DuplicateError

	if err := u.limiter.Wait(ctx); err != nil {
		return 0, entity.NewHTTPException(http.StatusInternalServerError, "")
	}

	if isValid := luhn.IsValidLuhn(orderID); !isValid {
		return 0, entity.NewHTTPException(http.StatusUnprocessableEntity, "")
	}

	userID, errContext := api.GetUserIDFromContext(ctx)

	if errContext != nil {
		return 0, entity.NewHTTPException(http.StatusUnauthorized, "")
	}

	createdOrderID, createOrderError := u.orderService.CreateOrder(ctx, orderID, userID)

	if createOrderError != nil {
		if errors.As(createOrderError, &duplicateError) {
			order, orderError := u.GetOrderByOrderID(ctx, orderID)

			if orderError != nil {
				return 0, entity.NewHTTPException(http.StatusInternalServerError, "")
			}

			var err *entity.HTTPException

			if order.UserID == userID {
				err = entity.NewHTTPException(http.StatusOK, "")
			} else {
				err = entity.NewHTTPException(http.StatusConflict, "")
			}

			return 0, err
		}

		return 0, createOrderError
	}

	go u.Accrual(ctx, orderID, userID)

	return createdOrderID, nil
}
