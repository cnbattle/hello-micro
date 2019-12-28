package jwt

import (
	"errors"
	"github.com/cnbattle/hello-micro/srv/user-srv/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// User model
type User struct {
	ID          uint     `json:"id"`
	Username    string   `json:"username" sql:"type:varchar(255), notnull, unique" binding:"required"`
	Password    string   `json:"-" sql:"type:varchar(255), notnull" binding:"required"`
	FullName    string   `json:"full_name" sql:"type:varchar(255)"`
	Permissions []string `json:"permissions"`
}

var (
	// Secret 加盐
	Secret = "hello-micro"
	// ExpireTime token有效期
	ExpireTime = 3600
)

// Claims token里面添加用户信息，验证token后可能会用到用户信息
type Claims struct {
	jwt.StandardClaims
	UserRole  int    `json:"user_role"`
	UID       string `json:"uid"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Gender    int    `json:"gender"`
	Signature string `json:"signature"`
}

// GenerateToken 生成 token
func GenerateToken(base *models.UserBase) (tokenStr string, err error) {
	var claims Claims
	claims.UserRole = base.UserRole
	claims.UID = base.UID
	claims.Nickname = base.Nickname
	claims.Avatar = base.Avatar
	claims.Gender = base.Gender
	claims.Signature = base.Signature
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Secret))
}

// ParseToken 解析验证token
func ParseToken(tokenSrt string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenSrt, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return &Claims{}, err
	}
	if err := token.Claims.Valid(); err != nil {
		return &Claims{}, err
	}
	if claim, ok := token.Claims.(*Claims); ok {
		return claim, nil
	}
	return &Claims{}, errors.New("Parse errors")
}

// VerifyToken 验证token
func VerifyToken(tokenSrt string) bool {
	token, err := jwt.ParseWithClaims(tokenSrt, &Claims{}, func(token *jwt.Token) (interface{}, error) {
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
