package model

import "time"

type Order struct {
	ID          uint          `gorm:"primaryKey"`
	UserID      uint          `gorm:"not null"`
	User        User          `gorm:"foreignKey:UserID"`
	Items       []OrderItem   `gorm:"foreignKey:OrderID"`
	Total       float64       `gorm:"not null"`
	Status      string        `gorm:"default:'pending'"`
	PaymentID   uint          `gorm:"not null"`
	Payment     *Payment      `gorm:"foreignKey:PaymentID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
