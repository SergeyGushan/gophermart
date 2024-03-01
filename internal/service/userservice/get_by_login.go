package userservice

import (
	"context"
	"gophermart/internal/entity"
)

func (s *Service) GetUserByLogin(ctx context.Context, login string) (entity.User, error) {
	return s.userRepo.GetUserByLogin(ctx, login)
}
