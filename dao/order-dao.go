package dao

type OrderDAO struct {
	CustomerName string `json:"customer_name" binding:"required"`
}
