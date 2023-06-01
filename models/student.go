package models

import "time"

type Student struct {
	ID          int64     `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Name        string    `json:"name"`
	Age         int64     `json:"age"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
