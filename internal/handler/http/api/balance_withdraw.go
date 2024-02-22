package api

import (
	"gophermart/internal/entity"
	"gophermart/internal/service"
	"net/http"
)

type RequestBalanceWithdraw struct {
	Order string  `json:"order"`
	Sum   float64 `json:"sum"`
}

func (h Handler) BalanceWithdraw(w http.ResponseWriter, r *http.Request) {
	var requestData RequestBalanceWithdraw

	if err := getDataFromRequest(r, &requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if isValid := service.IsValidLuhn(requestData.Order); !isValid {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	ctx := r.Context()

	userID, errContext := getUserIDFromContext(ctx)

	if errContext != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	balance, errGetBalanceByUserID := h.uc.GetBalanceByUserID(ctx, userID)

	if errGetBalanceByUserID != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if balance.Current < requestData.Sum {
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	_, errCreateOperation := h.uc.CreateOperation(
		ctx,
		requestData.Order,
		userID,
		entity.OperationTypeWithdrawn,
		requestData.Sum,
	)

	if errCreateOperation != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
