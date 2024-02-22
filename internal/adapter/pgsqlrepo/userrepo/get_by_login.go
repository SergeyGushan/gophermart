package userrepo

import (
	"context"
	"gophermart/internal/entity"
)

func (r *Repository) GetUserByLogin(ctx context.Context, login string) (entity.User, error) {
	user := entity.User{}

	err := r.Transactor.GetConn(ctx).QueryRow(
		"SELECT * FROM users WHERE login = $1",
		login,
	).Scan(&user.ID, &user.Login, &user.Password, &user.CreatedAt)

	return user, err
}
