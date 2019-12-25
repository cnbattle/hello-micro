// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	// 小程序登录
	MiniProgramLogin(ctx context.Context, in *MiniProgramLoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	// 手机验证码登录
	PhoneCodeLogin(ctx context.Context, in *PhoneCodeLoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	// 密码登录 (手机密码.用户名密码.邮箱密码)
	PhonePasswordLogin(ctx context.Context, in *PasswordLoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	AuthDetail(ctx context.Context, in *DefaultRequest, opts ...client.CallOption) (*AuthResponse, error)
	BaseDetail(ctx context.Context, in *DefaultRequest, opts ...client.CallOption) (*BaseResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) MiniProgramLogin(ctx context.Context, in *MiniProgramLoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "User.MiniProgramLogin", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) PhoneCodeLogin(ctx context.Context, in *PhoneCodeLoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "User.PhoneCodeLogin", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) PhonePasswordLogin(ctx context.Context, in *PasswordLoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "User.PhonePasswordLogin", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AuthDetail(ctx context.Context, in *DefaultRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "User.AuthDetail", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) BaseDetail(ctx context.Context, in *DefaultRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "User.BaseDetail", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	// 小程序登录
	MiniProgramLogin(context.Context, *MiniProgramLoginRequest, *LoginResponse) error
	// 手机验证码登录
	PhoneCodeLogin(context.Context, *PhoneCodeLoginRequest, *LoginResponse) error
	// 密码登录 (手机密码.用户名密码.邮箱密码)
	PhonePasswordLogin(context.Context, *PasswordLoginRequest, *LoginResponse) error
	AuthDetail(context.Context, *DefaultRequest, *AuthResponse) error
	BaseDetail(context.Context, *DefaultRequest, *BaseResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		MiniProgramLogin(ctx context.Context, in *MiniProgramLoginRequest, out *LoginResponse) error
		PhoneCodeLogin(ctx context.Context, in *PhoneCodeLoginRequest, out *LoginResponse) error
		PhonePasswordLogin(ctx context.Context, in *PasswordLoginRequest, out *LoginResponse) error
		AuthDetail(ctx context.Context, in *DefaultRequest, out *AuthResponse) error
		BaseDetail(ctx context.Context, in *DefaultRequest, out *BaseResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) MiniProgramLogin(ctx context.Context, in *MiniProgramLoginRequest, out *LoginResponse) error {
	return h.UserHandler.MiniProgramLogin(ctx, in, out)
}

func (h *userHandler) PhoneCodeLogin(ctx context.Context, in *PhoneCodeLoginRequest, out *LoginResponse) error {
	return h.UserHandler.PhoneCodeLogin(ctx, in, out)
}

func (h *userHandler) PhonePasswordLogin(ctx context.Context, in *PasswordLoginRequest, out *LoginResponse) error {
	return h.UserHandler.PhonePasswordLogin(ctx, in, out)
}

func (h *userHandler) AuthDetail(ctx context.Context, in *DefaultRequest, out *AuthResponse) error {
	return h.UserHandler.AuthDetail(ctx, in, out)
}

func (h *userHandler) BaseDetail(ctx context.Context, in *DefaultRequest, out *BaseResponse) error {
	return h.UserHandler.BaseDetail(ctx, in, out)
}