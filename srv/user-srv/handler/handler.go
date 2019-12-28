package handler

import (
	"context"
	user "github.com/cnbattle/hello-micro/srv/user-srv/proto"
)

// Handler struct
type Handler struct{}

// MiniProgramLogin 小程序登录
func (s Handler) MiniProgramLogin(ctx context.Context, req *user.MiniProgramLoginRequest, rsp *user.LoginResponse) error {

	return nil
}

// PhoneCodeLogin 手机登录
func (s Handler) PhoneCodeLogin(context.Context, *user.PhoneCodeLoginRequest, *user.LoginResponse) error {
	panic("implement me")
}

// PasswordLogin 密码登录
func (s Handler) PasswordLogin(context.Context, *user.PasswordLoginRequest, *user.LoginResponse) error {
	panic("implement me")
}

// AuthDetail 授权信息
func (s Handler) AuthDetail(context.Context, *user.DefaultRequest, *user.AuthResponse) error {
	panic("implement me")
}

// BaseDetail 基础信息
func (s Handler) BaseDetail(context.Context, *user.DefaultRequest, *user.BaseResponse) error {
	panic("implement me")
}
