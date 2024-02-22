package api

import (
	"encoding/json"
	"net/http"
)

func (h Handler) Balance(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusOK)
	balanceResourceJSON, _ := json.Marshal(balance)
	_, _ = w.Write(balanceResourceJSON)
}
