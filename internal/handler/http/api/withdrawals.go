package api

import (
	"encoding/json"
	"net/http"
)

func (h Handler) Withdrawals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, errContext := GetUserIDFromContext(ctx)

	if errContext != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	withdrawals, errWithdrawals := h.uc.GetOperationsByTypeWithdrawByUserID(ctx, userID)

	if errWithdrawals != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(withdrawals) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	listWithdrawalsJSON, _ := json.Marshal(withdrawals)
	_, _ = w.Write(listWithdrawalsJSON)
}
