package models

import (
	"time"

	"github.com/google/uuid"
)

// gorm.Model 的定义
type Model struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
