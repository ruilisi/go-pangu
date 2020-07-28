package args

type ChangePassword struct {
	OriginPassword  string `form:"origin_password" json:"origin_password" xml:"origin_password" binding:"required"`
	Password        string `form:"password" json:"password" xml:"password" binding:"required"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" xml:"password_confirm" binding:"required"`
}
type SignIn struct {
	Email    string `form:"email" json:"email" xml:"email" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	Device   string `form:"DEVICE_TYPE" json:"DEVICE_TYPE" xml:"DEVICE_TYPE" binding:"required"`
	Type     string `form:"type" json:"type" xml:"type" binding:"required"`
}
