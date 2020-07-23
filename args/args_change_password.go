package args

type ChangePassword struct {
	OriginPassword  string `form:"origin_password" json:"origin_password" xml:"origin_password" binding:"required"`
	Password        string `form:"password" json:"password" xml:"password" binding:"required"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" xml:"password_confirm" binding:"required"`
}
