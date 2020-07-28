package models

type User struct {
	Model
	Email             string `gorm:"index:idx_email,unique"`
	EncryptedPassword string
}
