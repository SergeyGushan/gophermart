package usecase

import (
	"context"
	"gophermart/internal/entity"
	"net/http"
)

func (u *UseCase) Accrual(ctx context.Context, orderID string, userID int64) {
	err := u.orderService.UpdateStatus(ctx, orderID, entity.OrderStatusProcessing)

	if err != nil {
		return
	}

	accrual, errAccrual := u.accrualService.GetOrderAccrual(orderID)

	if errAccrual != nil {
		return
	}

	if accrual.Code == http.StatusOK {
		if accrual.Item.Status == "INVALID" {
			_ = u.orderService.UpdateStatus(ctx, orderID, entity.OrderStatusInvalid)
			return
		}

		if accrual.Item.Status == "PROCESSED" {
			_ = u.orderService.UpdateStatusAndAccrual(ctx, orderID, entity.OrderStatusProcessed, accrual.Item.Accrual)
			_, _ = u.operationService.CreateOperation(ctx, orderID, userID, entity.OperationTypeAccrual, accrual.Item.Accrual)
		}
	}
}
