package model

import "time"



type User struct {
	ID        uint      `gorm:"primaryKey"`
	FirstName string    `gorm:"size:100;not null"`
	LastName  string    `gorm:"size:100;not null"`
	Email     string    `gorm:"uniqueIndex;size:100;not null"`
	Password  string    `gorm:"not null"`
	Phone     string    `gorm:"size:15"`
	Role      string    `gorm:"default:'customer';not null"` // customer o admin
	Address   []Address `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
