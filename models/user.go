package models

type User struct {
	Id                string
	Email             string `form:"email" json:"email" xml:"email" binding:"required"`
	Password          string `form:"password" json:"password" xml:"password" binding:"required"`
	EncryptedPassword string
	Device            string `form:"DEVICE_TYPE" json:"DEVICE_TYPE" xml:"DEVICE_TYPE" binding:"required"`
	Type              string `form:"type" json:"type" xml:"type" binding:"required"`
}
