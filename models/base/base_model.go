package base

import (
	"time"
)

type Model struct {
	ID        uint      `json:"id" gorm:"primarykey autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
