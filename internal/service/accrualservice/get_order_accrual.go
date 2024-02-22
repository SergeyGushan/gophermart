package accrualservice

import (
	"gophermart/internal/entity"
)

func (s *Service) GetOrderAccrual(orderID string) (entity.Accrual, error) {
	return s.accrualRepo.GetOrderAccrual(orderID)
}
