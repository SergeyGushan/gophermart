package api

import (
	"errors"
	"gophermart/internal/entity"
	"io"
	"net/http"
)

func (h Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var bodyReader io.Reader = r.Body
	var HTTPException *entity.HTTPException

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	body, err := io.ReadAll(bodyReader)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	orderID := string(body)

	_, createOrderError := h.uc.CreateOrder(ctx, orderID)
	if createOrderError != nil {
		if errors.As(createOrderError, &HTTPException) {
			w.WriteHeader(HTTPException.Code)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}

		return
	}

	w.WriteHeader(http.StatusAccepted)
}
