package main

import (
	"context"

	"github.com/cnbattle/hello-micro/proto/login"
)

type handler struct {
}

func (h handler) MiniProgramLogin(context.Context, *login.MiniProgramLoginRequest, *login.LoginResponse) error {
	panic("implement me")
}

func (h handler) PhoneCodeLogin(context.Context, *login.PhoneCodeLoginRequest, *login.LoginResponse) error {
	panic("implement me")
}

func (h handler) EmailCodeLogin(context.Context, *login.EmailCodeLoginRequest, *login.LoginResponse) error {
	panic("implement me")
}

func (h handler) PhonePasswordLogin(context.Context, *login.PasswordLoginRequest, *login.LoginResponse) error {
	panic("implement me")
}
