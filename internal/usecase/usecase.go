package usecase

import (
	"context"
	"golang.org/x/time/rate"
	"gophermart/internal/entity"
)

type userService interface {
	CreateUser(ctx context.Context, login string, password string) (entity.User, error)
	GetUserByLogin(ctx context.Context, login string) (entity.User, error)
}

type orderService interface {
	CreateOrder(ctx context.Context, orderID string, userID int64) (int64, error)
	GetByOrderID(ctx context.Context, orderID string) (entity.Order, error)
	GetByUserID(ctx context.Context, userID int64) ([]entity.Order, error)
	UpdateStatus(ctx context.Context, orderID string, status entity.OrderStatus) error
	UpdateStatusAndAccrual(ctx context.Context, orderID string, status entity.OrderStatus, accrual float64) error
}

type operationService interface {
	GetBalanceByUserID(ctx context.Context, userID int64) (entity.BalanceResponse, error)
	CreateOperation(ctx context.Context, orderID string, userID int64, operationType entity.OperationType, sum float64) (int64, error)
	GetOperationsByTypeWithdrawByUserID(ctx context.Context, userID int64) ([]entity.Operation, error)
}

type accrualService interface {
	GetOrderAccrual(orderID string) (entity.Accrual, error)
}

type UseCase struct {
	userService      userService
	orderService     orderService
	operationService operationService
	accrualService   accrualService
	limiter          *rate.Limiter
}

func NewUseCase(userService userService, orderService orderService, operationService operationService, accrualService accrualService) *UseCase {
	return &UseCase{
		userService:      userService,
		orderService:     orderService,
		operationService: operationService,
		accrualService:   accrualService,
		limiter:          rate.NewLimiter(rate.Limit(10), 1),
	}
}
