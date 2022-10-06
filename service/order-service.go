package service

import (
	"github.com/husfuu/go-order/dao"
	"github.com/husfuu/go-order/entity"
	"github.com/husfuu/go-order/repository"
)

type OrderService interface {
	GetOrders() ([]entity.Order, error)
	GetOrderById(input dao.GetOrderDetailInput) (entity.Order, error)
	CreateOrder(input dao.CreateOrderInput) (entity.Order, error)
	// UpdateOrder(inputId dao.GetOrderDetailInput, inputData dao.CreateOrderInput) (entity.Order, error)
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

func (s *orderService) GetOrderById(input dao.GetOrderDetailInput) (entity.Order, error) {
	order, err := s.repository.FindById(input.ID)

	if err != nil {
		return order, nil
	}

	return order, nil
}

func (s *orderService) CreateOrder(input dao.CreateOrderInput) (entity.Order, error) {
	order := entity.Order{}
	order.CustomerName = input.CustomerName

	var items = make([]entity.Item, 0)
	items = append(items, input.OrderItems...)
	order.OrderItems = items

	newOrder, err := s.repository.Save(order)

	if err != nil {
		return newOrder, nil
	}

	return newOrder, nil
}
