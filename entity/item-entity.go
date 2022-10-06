package entity

import (
	"time"
)

type Item struct {
	ID          int       `json:"id" gorm:"primary_key"`
	OrderID     int       `json:"order_id" gorm:"order_id"`
	Code        string    `json:"code" gorm:"column:code;"`
	Description string    `json:"description" gorm:"column:description;"`
	Quantity    int       `json:"quantity" gorm:"column:quantity;"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"column:deleted_at;"`
}
