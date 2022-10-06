package entity

import "time"

type Item struct {
	ID          int
	OrderId     int
	Code        string
	Description string
	Quantity    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
