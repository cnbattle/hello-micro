package handler

import (
	"context"
	"errors"
	loginservice "github.com/cnbattle/hello-micro/srv/login-srv/service"

	"github.com/cnbattle/hello-micro/proto/login"
)

type Handler struct {
}

// MiniProgramLogin 小程序登录
func (h Handler) MiniProgramLogin(ctx context.Context, req *login.MiniProgramLoginRequest, rsp *login.LoginResponse) error {
	miniProgramLoginService := loginservice.MiniProgramLoginService{
		Channel:    req.Channel,
		Code:       req.Code,
		AutoCreate: true,
	}
	request, err := miniProgramLoginService.Login()
	if err != nil {
		return err
	}
	rsp.Token = request.Token
	return nil
}

func (h Handler) PhoneCodeLogin(context.Context, *login.PhoneCodeLoginRequest, *login.LoginResponse) error {
	return errors.New("暂不支持")
}

func (h Handler) EmailCodeLogin(context.Context, *login.EmailCodeLoginRequest, *login.LoginResponse) error {
	return errors.New("暂不支持")
}

// PasswordLogin 密码登录
func (h Handler) PasswordLogin(ctx context.Context, req *login.PasswordLoginRequest, rsp *login.LoginResponse) error {
	passwordLoginService := loginservice.PasswordLoginService{
		Username: req.LoginName,
		Password: req.Password,
	}
	request, err := passwordLoginService.Login()
	if err != nil {
		return err
	}
	rsp.Token = request.Token
	return nil
}
