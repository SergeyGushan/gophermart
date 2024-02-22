package api

import (
	"errors"
	"gophermart/internal/entity"
	"gophermart/internal/service"
	"net/http"
)

type RequestCreateUserData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var duplicateError *entity.DuplicateError
	var requestData RequestCreateUserData

	if err := getDataFromRequest(r, &requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	user, createUserError := h.uc.CreateUser(ctx, requestData.Login, requestData.Password)

	if errors.As(createUserError, &duplicateError) {
		w.WriteHeader(http.StatusConflict)
		return
	}

	if createUserError != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, setTokenToCookieError := service.SetTokenToCookie(w, user)
	if setTokenToCookieError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
