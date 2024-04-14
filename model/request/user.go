package request

type Login struct {
	Account  string `json:"account" binding:"required" example:"0912345678"` // 帳號
	Password string `json:"password" binding:"required" example:"123456"`    // 密碼
	// Captcha   string `json:"captcha" binding:"required"`   // 驗證碼
}

type Register struct {
	Username  string `json:"userName" binding:"required" example:"Edward"`
	Password  string `json:"passWord" binding:"required" example:"123456"`
	NickName  string `json:"nickName" example:"Ed"`
	HeaderImg string `json:"headerImg" example:"http://headimgurl.com"`
	Phone     string `json:"phone" binding:"required" example:"0912345678"`
	Email     string `json:"email" binding:"required" example:"123@gmail.com"`
}
