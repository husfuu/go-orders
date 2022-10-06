package entity

import "time"

type Order struct {
	ID           int
	CustomerName string
	OrderItems   []Item
	OrderedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
