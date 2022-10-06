package entity

import (
	"time"
)

type Order struct {
	ID           int       `json:"id" gorm:"primary_key"`
	CustomerName string    `json:"customer_name" gorm:"column:customer_name;"`
	OrderItems   []Item    `json:"order_items" gorm:"foreignkey:OrderID;association_foreignkey:OrderID"`
	OrderedAt    time.Time `json:"ordered_at" gorm:"column:ordered_at;"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at;"`
	DeletedAt    time.Time `json:"deleted_at" gorm:"column:deleted_at;"`
}

type CreateOrder struct {
	ID           int
	CustomerName string
	OrderIds     []int
	OrderedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
