package userrepo

import (
	"context"
	"gophermart/internal/entity"
)

func (r *Repository) GetUserByID(ctx context.Context, userID int64) (entity.User, error) {
	user := entity.User{}

	err := r.Transactor.GetConn(ctx).QueryRow(
		"SELECT * FROM users WHERE id = $1",
		userID,
	).Scan(&user.ID, &user.Login, &user.Password, &user.CreatedAt)

	return user, err
}
