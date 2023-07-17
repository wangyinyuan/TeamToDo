package request

// UserRegisterRequest 注册时的请求
type UserRegisterRequest struct {
	Email    string `json:"email" validate:"required,email" `
	UserName string `json:"userName" validate:"required,min=3,max=20" `
	Password string `json:"password" validate:"required,min=8,max=20" `
}

// UserSignInRequest 登陆时的请求
type UserSignInRequest struct {
	Email    string `json:"email" validate:"required,email" `
	Password string `json:"password" validate:"required,min=8,max=20" `
}

// UserUpdateRequest 更新用户信息时的请求(包括密码)
type UserUpdateRequest struct {
	UserID         uint   `json:"-"`
	UserName       string `validate:"omitempty,min=3,max=20" json:"userName,omitempty"`
	OriginPassword string `validate:"omitempty,min=8,max=20" json:"originPassword,omitempty"`
	Password       string `validate:"omitempty,min=8,max=20" json:"password,omitempty"`
	Email          string `validate:"omitempty,email" json:"email,omitempty"`
}
