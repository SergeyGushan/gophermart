package api

import (
	"context"
	"gophermart/internal/entity"
	"gophermart/pkg/logger"
)

type UseCase interface {
	CreateUser(ctx context.Context, login string, password string) (entity.User, error)
	GetUserByLogin(ctx context.Context, login string) (entity.User, error)
	CreateOrder(ctx context.Context, orderID string, userID int64) (int64, error)
	GetOrderByOrderID(ctx context.Context, orderID string) (entity.Order, error)
	GetOrdersByUserID(ctx context.Context, userID int64) ([]entity.Order, error)
	GetOperationsByTypeWithdrawByUserID(ctx context.Context, userID int64) ([]entity.Operation, error)
	CreateOperation(ctx context.Context, orderID string, userID int64, operationType entity.OperationType, sum float64) (int64, error)
	GetBalanceByUserID(ctx context.Context, userID int64) (entity.BalanceResponse, error)
	Accrual(ctx context.Context, orderID string, userID int64)
}

type Handler struct {
	uc     UseCase
	logger logger.Logger
}

func NewHandler(uc UseCase, logs logger.Logger) *Handler {
	return &Handler{uc: uc, logger: logs}
}
