package userrepo

import (
	"context"
	"gophermart/internal/entity"
)

func (r *Repository) CreateUser(ctx context.Context, login string, password string) (int64, error) {
	var lastInsertID int64

	err := r.Transactor.GetConn(ctx).QueryRow(
		"INSERT INTO users (login, password) VALUES ($1, $2) ON CONFLICT (login) DO NOTHING RETURNING id",
		login, password,
	).Scan(&lastInsertID)

	if lastInsertID == 0 {
		return lastInsertID, &entity.DuplicateError{
			Field: "login",
			Value: login,
		}
	}

	if err != nil {
		return lastInsertID, err
	}

	return lastInsertID, nil
}
