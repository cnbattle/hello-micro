package loginservice

import (
	"errors"
	"github.com/cnbattle/hello-micro/pkg/jwt"
	"github.com/cnbattle/hello-micro/srv/user-srv/models"
	"strings"

	"github.com/cnbattle/hello-micro/pkg/config"
	"github.com/cnbattle/hello-micro/pkg/database"
	"github.com/silenceper/wechat"
)

// MiniProgramLoginService MiniProgramLoginService
type MiniProgramLoginService struct {
	Channel    string
	Code       string
	AutoCreate bool
}

// Login Login
func (m *MiniProgramLoginService) Login() (LoginRequest, error) {
	if m.Code == "" {
		return LoginRequest{}, errors.New("code is empty")
	}
	wxConfig := &wechat.Config{
		AppID:     config.GetEnv("WEAPP_" + strings.ToUpper(m.Channel) + "_APP_ID"),
		AppSecret: config.GetEnv("WEAPP_" + strings.ToUpper(m.Channel) + "_APP_SECRET"),
	}
	session, err := wechat.NewWechat(wxConfig).GetMiniProgram().Code2Session(m.Code)
	if err != nil {
		return LoginRequest{}, errors.New("Code2Session is error:" + err.Error())
	}
	// 用户登录
	userAuth, err := loginOrCreate(m.Channel, session.OpenID, m.AutoCreate)
	if err != nil {
		return LoginRequest{}, err
	}
	// 判断是否有基础用户信息
	var userBase models.UserBase
	err = database.Conn.Where("uid=?", userAuth.UID).Find(&userBase).Error
	if err != nil {
		return LoginRequest{}, err
	}
	tokenStr, err := jwt.GenerateToken(&userBase)
	if err != nil {
		return LoginRequest{}, err
	}
	return LoginRequest{Token: tokenStr}, err
}
