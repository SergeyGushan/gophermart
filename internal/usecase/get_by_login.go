package usecase

import (
	"context"
	"gophermart/internal/entity"
)

func (u *UseCase) GetUserByLogin(ctx context.Context, login string) (entity.User, error) {
	return u.userService.GetUserByLogin(ctx, login)
}
