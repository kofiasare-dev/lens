package models

import (
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // Auto-generated created timestamp
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // Auto-generated updated timestamp
}
