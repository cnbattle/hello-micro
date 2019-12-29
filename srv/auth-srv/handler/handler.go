package handler

import (
	"context"
	"errors"

	"github.com/cnbattle/hello-micro/proto/auth"
	"github.com/cnbattle/hello-micro/srv/auth-srv/service"
)

// Handler auth proto struct
type Handler struct {
}

// VerifyToken 验证token
func (h Handler) VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest, rsp *auth.VerifyTokenResponse) error {
	rsp.Valid = service.VerifyToken(req.Token)
	return nil
}

// RefreshToken 刷新token
func (h Handler) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest, rsp *auth.RefreshTokenResponse) error {
	token, err := service.RefreshToken(req.Token)
	if err != nil {
		return err
	}
	rsp.Token = token
	return nil
}

// MiniProgramLogin 小程序登录
func (h Handler) MiniProgramLogin(ctx context.Context, req *auth.MiniProgramLoginRequest, rsp *auth.LoginResponse) error {
	miniProgramLoginService := service.MiniProgramLoginService{
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

// PhoneCodeLogin 手机登录
func (h Handler) PhoneCodeLogin(context.Context, *auth.PhoneCodeLoginRequest, *auth.LoginResponse) error {
	return errors.New("暂不支持")
}

// EmailCodeLogin 邮箱登录
func (h Handler) EmailCodeLogin(context.Context, *auth.EmailCodeLoginRequest, *auth.LoginResponse) error {
	return errors.New("暂不支持")
}

// PasswordLogin 密码登录
func (h Handler) PasswordLogin(ctx context.Context, req *auth.PasswordLoginRequest, rsp *auth.LoginResponse) error {
	passwordLoginService := service.PasswordLoginService{
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
