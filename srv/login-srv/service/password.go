package loginservice

import (
	"errors"

	"github.com/cnbattle/hello-micro/pkg/database"
	"github.com/cnbattle/hello-micro/pkg/utils"
	"github.com/cnbattle/hello-micro/srv/user-srv/models"
)

type PasswordLoginService struct {
	Username   string
	Password   string
	AutoCreate bool
}

func (m *PasswordLoginService) Login() (LoginRequest, error) {
	if m.Password == "" || m.Username == "" {
		return LoginRequest{}, errors.New("code is empty")
	}
	// 用户登录
	userAuth, err := loginOrCreate(m.Username, utils.MD5Sum(m.Password), m.AutoCreate)
	if err != nil {
		return LoginRequest{}, err
	}
	// 判断是否有基础用户信息
	var userBase models.UserBase
	err = database.Conn.Where("uid=?", userAuth.Uid).Find(&userBase).Error
	if err != nil {
		return LoginRequest{Uid: userAuth.Uid, Token: userAuth.Certificate}, err
	}
	return LoginRequest{Uid: userAuth.Uid, Token: userAuth.Certificate, IsAuth: true}, err
}
