package service

import "github.com/cnbattle/hello-micro/pkg/jwt"

// VerifyToken 验证token
func VerifyToken(token string) bool {
	return jwt.VerifyToken(token)
}
