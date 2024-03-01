package api

import (
	"gophermart/internal/service"
	"net/http"
)

type RequestAuthUserData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	var requestData RequestAuthUserData

	if err := getDataFromRequest(r, &requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	user, err := h.uc.GetUserByLogin(ctx, requestData.Login)

	if err != nil || !service.ComparePasswords(user.Password, requestData.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, setTokenToCookieError := service.SetTokenToCookie(w, user)
	if setTokenToCookieError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
