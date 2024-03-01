package api

import (
	"encoding/json"
	"errors"
	"gophermart/internal/entity"
	"net/http"
)

func (h Handler) Balance(w http.ResponseWriter, r *http.Request) {
	var HTTPException *entity.HTTPException

	ctx := r.Context()
	balance, errGetBalanceByUserID := h.uc.GetBalanceByUserID(ctx)

	if errGetBalanceByUserID != nil {
		if errors.As(errGetBalanceByUserID, &HTTPException) {
			w.WriteHeader(HTTPException.Code)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	balanceResourceJSON, _ := json.Marshal(balance)
	_, _ = w.Write(balanceResourceJSON)
}
