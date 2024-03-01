package service

import (
	"gophermart/internal/authentication"
	"gophermart/internal/entity"
	"net/http"
	"time"
)

func SetTokenToCookie(res http.ResponseWriter, user entity.User) (string, error) {
	token, err := authentication.BuildJWTString(user.ID)

	if err == nil {
		http.SetCookie(res, &http.Cookie{
			Name:    authentication.TokenKey,
			Value:   token,
			Expires: time.Now().Add(authentication.TokenExp),
			Path:    "/",
		})
	}

	return token, err
}
