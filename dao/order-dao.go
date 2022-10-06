package dao

type CreateOrderInput struct {
	CustomerName string `json:"customer_name" binding:"required"`
}

type GetOrderDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
