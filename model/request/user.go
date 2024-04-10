package request

type Login struct {
	Account  string `json:"account" binding:"required"`  // 帳號
	Password  string `json:"password" binding:"required"`  // 密碼
	// Captcha   string `json:"captcha" binding:"required"`   // 驗證碼
}

type Register struct {
	Username     string `json:"userName" example:"使用者名稱"`
	Password     string `json:"passWord" example:"密碼"`
	NickName     string `json:"nickName" example:"暱稱"`
	HeaderImg    string `json:"headerImg" example:"頭像連結"`
	Phone        string `json:"phone" example:"電話號碼"`
	Email        string `json:"email" example:"電子信箱"`
}