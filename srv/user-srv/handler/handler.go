package handler

import (
	"context"
	"github.com/cnbattle/hello-micro/proto/user"
)

type Handler struct{}

func (s Handler) Login(context.Context, *user.LoginRequest, *user.AuthResponse) error {
	panic("implement me")
}

func (s Handler) AuthDetail(context.Context, *user.AuthRequest, *user.AuthResponse) error {
	panic("implement me")
}

func (s Handler) BaseDetail(context.Context, *user.BaseRequest, *user.BaseResponse) error {
	panic("implement me")
}
