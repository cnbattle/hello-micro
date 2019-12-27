package handler

import (
	"context"
	user "github.com/cnbattle/hello-micro/srv/user-srv/proto"
)

type Handler struct{}

// MiniProgramLogin 小程序登录
func (s Handler) MiniProgramLogin(ctx context.Context, req *user.MiniProgramLoginRequest, rsp *user.LoginResponse) error {

	return nil
}

func (s Handler) PhoneCodeLogin(context.Context, *user.PhoneCodeLoginRequest, *user.LoginResponse) error {
	panic("implement me")
}

func (s Handler) PasswordLogin(context.Context, *user.PasswordLoginRequest, *user.LoginResponse) error {
	panic("implement me")
}

func (s Handler) AuthDetail(context.Context, *user.DefaultRequest, *user.AuthResponse) error {
	panic("implement me")
}

func (s Handler) BaseDetail(context.Context, *user.DefaultRequest, *user.BaseResponse) error {
	panic("implement me")
}
