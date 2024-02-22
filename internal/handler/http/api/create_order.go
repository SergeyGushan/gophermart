package api

import (
	"errors"
	"gophermart/internal/entity"
	"gophermart/internal/service"
	"io"
	"net/http"
)

func (h Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var bodyReader io.Reader = r.Body
	var duplicateError *entity.DuplicateError

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	body, err := io.ReadAll(bodyReader)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	orderID := string(body)

	if isValid := service.IsValidLuhn(orderID); !isValid {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	ctx := r.Context()

	userID, errContext := getUserIDFromContext(ctx)

	if errContext != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, saveOrderError := h.uc.CreateOrder(ctx, orderID, userID)

	if errors.As(saveOrderError, &duplicateError) {
		var orderError error
		order, orderError := h.uc.GetOrderByOrderID(ctx, orderID)

		if orderError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if order.UserID == userID {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusConflict)
		}

		return
	}

	if saveOrderError != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	go h.uc.Accrual(ctx, orderID, userID)

	w.WriteHeader(http.StatusAccepted)
}
