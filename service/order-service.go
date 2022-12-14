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
	UpdateOrder(inputId dao.GetOrderDetailInput, inputData dao.CreateOrderInput) (entity.Order, error)
	DeleteOrder(inputId dao.GetOrderDetailInput) (entity.Order, error)
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
	// fmt.Println(order)
	if err != nil {
		return order, err
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

func (s *orderService) UpdateOrder(inputID dao.GetOrderDetailInput, inputData dao.CreateOrderInput) (entity.Order, error) {
	order, err := s.repository.FindById(inputID.ID)

	if order.ID == 0 {
		return order, err
	}

	if err != nil {
		return order, err
	}

	order.CustomerName = inputData.CustomerName

	order.OrderItems = []entity.Item{}

	var items = make([]entity.Item, 0)
	items = append(items, inputData.OrderItems...)
	order.OrderItems = items
	updatedOrder, err := s.repository.Update(order)

	if err != nil {
		return updatedOrder, nil
	}

	return updatedOrder, nil
}

func (s *orderService) DeleteOrder(inputID dao.GetOrderDetailInput) (entity.Order, error) {
	order, err := s.repository.FindById(inputID.ID)

	if err != nil {
		return order, err
	}

	s.repository.Delete(inputID.ID)

	return order, nil
}
