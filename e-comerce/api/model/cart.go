package model

import "time"

type Cart struct {
	ID        uint         `gorm:"primaryKey"`
	UserID    uint         `gorm:"not null"`
	User      User         `gorm:"foreignKey:UserID"`
	Items     []CartItem   `gorm:"foreignKey:CartID"`
	Total     float64      `gorm:"not null"`
	UpdatedAt time.Time
}
