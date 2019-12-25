package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// model
type User struct {
	Id          uint     `json:"id"`
	Username    string   `json:"username" sql:"type:varchar(255), notnull, unique" binding:"required"`
	Password    string   `json:"-" sql:"type:varchar(255), notnull" binding:"required"`
	FullName    string   `json:"full_name" sql:"type:varchar(255)"`
	Permissions []string `json:"permissions"`
}

var (
	Secret     = "test_secret" // 加盐
	ExpireTime = 10            // token有效期
)

type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	Uid         string   `json:"uid"`
	Nickname    string   `json:"nickname"`
	Avatar      string   `json:"avatar"`
	Gender      string   `json:"gender"`
	Signature   string   `json:"signature"`
	Permissions []string `json:"permissions"`
}

// GenerateToken 生成 token
func GenerateToken(uid, nickname, avatar, gender, signature string) (tokenStr string, err error) {
	var claims JWTClaims
	claims.Uid = uid
	claims.Nickname = nickname
	claims.Avatar = avatar
	claims.Gender = gender
	claims.Signature = signature
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Secret))
}

// 解析验证token
func ParseToken(tokenSrt string) (claims *JWTClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenSrt, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return &JWTClaims{}, err
	}
	if err := token.Claims.Valid(); err != nil {
		return &JWTClaims{}, err
	}
	if claim, ok := token.Claims.(*JWTClaims); ok {
		return claim, nil
	} else {
		return &JWTClaims{}, errors.New("Parse errors")
	}
}

// 验证token
func VerifyToken(tokenSrt string) bool {
	token, err := jwt.ParseWithClaims(tokenSrt, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return false
	}
	if err := token.Claims.Valid(); err != nil {
		return false
	}
	return true
}
