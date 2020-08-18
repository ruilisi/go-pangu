package models

import (
	"go-pangu/db"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// gorm.Model 的定义
type Model struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Seed() {
	db.Open()
	email := "test@123.com"
	user := FindUserByEmail(email)
	if user.Email == "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte("test123"), bcrypt.DefaultCost)
		user = &User{Email: email, EncryptedPassword: string(hash)}
		db.Create()
	}
}
