package service

import (
	"github.com/husfuu/go-order/entity"
	"github.com/husfuu/go-order/repository"
)

type OrderService interface {
	GetOrders() ([]entity.Order, error)
}

type orderService struct {
	repository repository.OrderRepository
}

func NewOrderService(repository repository.OrderRepository) *orderService {
	return &orderService{repository}
}

func (s *orderService) GetOrders() ([]entity.Order, error) {
	orders, err := s.repository.FindAll()

	if err != nil {
		return orders, nil
	}

	return orders, nil
}
