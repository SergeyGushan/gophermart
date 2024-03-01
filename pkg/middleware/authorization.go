package middleware

import (
	"context"
	"errors"
	"gophermart/internal/authentication"
	"gophermart/internal/entity"
	"net/http"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		token, tokenErr := getTokenFromCookie(req)
		var TokenError *entity.TokenError

		if errors.As(tokenErr, &TokenError) {
			res.WriteHeader(http.StatusUnauthorized)
			return
		}

		userID, uError := getUserIDByToken(token)

		if uError != nil {
			res.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), UserIDKey, userID)
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	})
}

func getUserIDByToken(token string) (int64, error) {
	return authentication.GetUserIDFromJWTString(token)
}

func getTokenFromCookie(req *http.Request) (string, error) {
	tokenCookie, err := req.Cookie(authentication.TokenKey)
	if err != nil {
		return "", err
	}

	token := tokenCookie.Value

	return token, nil
}
