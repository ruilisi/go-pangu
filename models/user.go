package models

type User struct {
	Model
	Email             string `gorm:"index:idx_email,unique"`
	EncryptedPassword string
}

func FindUserByEmail(email string) *User {
	var user User
	DB.Where("email = ?", email).First(&user)
	return &user
}

func FindUserById(id string) *User {
	var user User
	DB.Where("id = ?", id).First(&user)
	return &user
}
