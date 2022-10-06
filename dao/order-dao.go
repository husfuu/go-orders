package dao

import "github.com/husfuu/go-order/entity"

type CreateOrderInput struct {
	CustomerName string        `json:"customer_name" binding:"required"`
	OrderItems   []entity.Item `json:"order_items" binding:"required"`
}

type GetOrderDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
