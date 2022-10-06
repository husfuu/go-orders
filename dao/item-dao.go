package dao

import "time"

type ItemDAO struct {
	OrderId     int    `json:"order_id" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
