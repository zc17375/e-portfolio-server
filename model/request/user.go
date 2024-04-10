package request

type Login struct {
	Account  string `json:"account" binding:"required" example:"信箱或密碼"`  // 帳號
	Password  string `json:"password" binding:"required" example:"密碼"`  // 密碼
	// Captcha   string `json:"captcha" binding:"required"`   // 驗證碼
}

type Register struct {
	Username     string `json:"userName" binding:"required" example:"使用者名稱"`
	Password     string `json:"passWord" binding:"required" example:"密碼"`
	NickName     string `json:"nickName" example:"暱稱"`
	HeaderImg    string `json:"headerImg" example:"頭像連結"`
	Phone        string `json:"phone" binding:"required" example:"電話號碼"`
	Email        string `json:"email" binding:"required" example:"電子信箱"`
}