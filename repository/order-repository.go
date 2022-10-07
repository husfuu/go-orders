package repository

import (
	"github.com/husfuu/go-order/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll() ([]entity.Order, error)
	FindById(ID int) (entity.Order, error)
	Save(order entity.Order) (entity.Order, error)
	Update(order entity.Order) (entity.Order, error)
	Delete(ID int) (entity.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) FindAll() ([]entity.Order, error) {
	var orders []entity.Order

	err := r.db.Preload("OrderItems").Find(&orders).Error

	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *orderRepository) FindById(ID int) (entity.Order, error) {
	var order entity.Order

	err := r.db.Preload("OrderItems").Where("id = ?", ID).Find(&order).Error

	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *orderRepository) Save(order entity.Order) (entity.Order, error) {
	err := r.db.Create(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *orderRepository) Update(order entity.Order) (entity.Order, error) {

	r.DeleteItems(order.ID)

	err := r.db.Save(&order)

	if err != nil {
		return order, nil
	}

	return order, nil
}

func (r *orderRepository) DeleteItems(ID int) {
	r.db.Where("order_id = ?", ID).Delete(&entity.Item{})
}

func (r *orderRepository) Delete(ID int) (entity.Order, error) {
	var order entity.Order

	r.DeleteItems(ID)

	err := r.db.Where("id = ?", ID).Delete(&order).Error

	if err != nil {
		return order, err
	}
	return order, nil
}
