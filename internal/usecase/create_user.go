package usecase

import (
	"context"
	"gophermart/internal/entity"
)

func (u *UseCase) CreateUser(ctx context.Context, login string, password string) (entity.User, error) {
	return u.userService.CreateUser(ctx, login, password)
}
