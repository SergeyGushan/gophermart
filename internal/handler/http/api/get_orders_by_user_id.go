package api

import (
	"encoding/json"
	"net/http"
)

func (h Handler) GetOrdersByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, errContext := GetUserIDFromContext(ctx)

	if errContext != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	orders, err := h.uc.GetOrdersByUserID(ctx, userID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(orders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	listOrdersJSON, _ := json.Marshal(orders)
	_, _ = w.Write(listOrdersJSON)
}
