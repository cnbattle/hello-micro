package handler

import (
	"context"
	"github.com/cnbattle/hello-micro/srv/user-srv/service"

	user "github.com/cnbattle/hello-micro/srv/user-srv/proto"
)

type Handler struct{}

// MiniProgramLogin 小程序登录
func (s Handler) MiniProgramLogin(ctx context.Context, req *user.MiniProgramLoginRequest, rsp *user.LoginResponse) error {
	token, err := service.MiniProgramLogin(req.Channel, req.Code)
	if err != nil {
		return err
	}
	rsp.Token = token
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
