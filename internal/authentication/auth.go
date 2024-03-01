package authentication

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"gophermart/internal/entity"
	"time"
)

const SecretKey = "superSecretKey"
const TokenExp = time.Hour * 24
const TokenKey = "token"

type Claims struct {
	jwt.RegisteredClaims
	UserID int64
}

func BuildJWTString(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExp)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserIDFromJWTString(tokenString string) (int64, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return 0, entity.TokenError{}
	}

	if !token.Valid {
		return 0, entity.TokenError{}
	}

	return claims.UserID, nil
}
