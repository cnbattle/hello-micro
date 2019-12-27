package service

import (
	"errors"
	"github.com/cnbattle/hello-micro/pkg/jwt"
	"strings"

	"github.com/cnbattle/hello-micro/pkg/config"
	"github.com/cnbattle/hello-micro/srv/user-srv/database"
	"github.com/cnbattle/hello-micro/srv/user-srv/models"
	"github.com/google/uuid"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/util"
)

// MiniProgramLogin 小程序登录
func MiniProgramLogin(channel, code string) (token string, err error) {
	if code == "" {
		return "", errors.New("code is empty")
	}
	wxConfig := &wechat.Config{
		AppID:     config.GetEnv("WEAPP_" + strings.ToUpper(channel) + "_APP_ID"),
		AppSecret: config.GetEnv("WEAPP_" + strings.ToUpper(channel) + "_APP_SECRET"),
	}
	session, err := wechat.NewWechat(wxConfig).GetMiniProgram().Code2Session(code)
	if err != nil {
		return "", errors.New("Code2Session is error:" + err.Error())
	}
	// 用户登录
	userAuth, err := LoginOrCreate(channel, session.OpenID, true)
	if err != nil {
		return "", err
	}
	// 判断是否有基础用户信息
	var userBase models.UserBase
	err = database.UserConn.Where("uid=?", userAuth.Uid).Find(&userBase).Error
	if err != nil {
		return "", err
	}
	tokenStr, err := jwt.GenerateToken(&userBase)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func AuthDetail(uid string) {

}

func BaseDetail(uid string) {

}

// LoginOrCreate
func LoginOrCreate(identityType, identity string, autoCreate bool) (*models.UserAuth, error) {
	userAuth := models.UserAuth{}
	err := database.UserConn.Where("identity_type=?", identityType).
		Where("identity=?", identity).First(&userAuth).Error
	if err != nil {
		if autoCreate {
			userAuth.IdentityType = identityType
			userAuth.Identity = identity
			userAuth.Uid = uuid.New().String()
			database.UserConn.Create(&userAuth)
		} else {
			return &models.UserAuth{}, errors.New("not has user auth data")
		}
	}
	userAuth.Certificate = util.MD5Sum(uuid.New().String())
	database.UserConn.Save(&userAuth)
	return &userAuth, nil
}
