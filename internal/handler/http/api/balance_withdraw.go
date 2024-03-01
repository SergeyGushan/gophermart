package api

import (
	"errors"
	"gophermart/internal/entity"
	"net/http"
)

type RequestBalanceWithdraw struct {
	Order string  `json:"order"`
	Sum   float64 `json:"sum"`
}

func (h Handler) BalanceWithdraw(w http.ResponseWriter, r *http.Request) {
	var requestData RequestBalanceWithdraw
	var HTTPException *entity.HTTPException

	if err := getDataFromRequest(r, &requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	_, errCreateOperation := h.uc.CreateWithdrawnOperation(ctx, requestData.Order, requestData.Sum)

	if errCreateOperation != nil {
		if errors.As(errCreateOperation, &HTTPException) {
			w.WriteHeader(HTTPException.Code)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
}
