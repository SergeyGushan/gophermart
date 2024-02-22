package userservice

import (
	"context"
	"gophermart/internal/entity"
	"gophermart/internal/service"
)

func (s *Service) CreateUser(ctx context.Context, login string, password string) (entity.User, error) {
	encodedPass, errEncodePass := service.EncodePassword(password)

	if errEncodePass != nil {
		return entity.User{}, errEncodePass
	}

	userID, errCreateUser := s.userRepo.CreateUser(ctx, login, string(encodedPass))

	if errCreateUser != nil {
		return entity.User{}, errCreateUser
	}

	return s.userRepo.GetUserByID(ctx, userID)
}
