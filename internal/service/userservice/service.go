package userservice

import (
	"context"
	"gophermart/internal/entity"
	"gophermart/internal/service"
)

type userRepo interface {
	service.Transactor
	CreateUser(ctx context.Context, login string, password string) (int64, error)
	GetUserByID(ctx context.Context, userID int64) (entity.User, error)
	GetUserByLogin(ctx context.Context, login string) (entity.User, error)
}

type Service struct {
	userRepo userRepo
}

func NewService(userRepo userRepo) *Service {
	return &Service{
		userRepo: userRepo,
	}
}
