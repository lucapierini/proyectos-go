package model

import "time"

type Payment struct {
	ID          uint      `gorm:"primaryKey"`
	OrderID     uint      `gorm:"not null"`
	Order       Order     `gorm:"foreignKey:OrderID"`
	Method      string    `gorm:"size:50;not null"`
	Status      string    `gorm:"default:'pending'"`
	Transaction string    `gorm:"size:100"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
