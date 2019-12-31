package service

import "github.com/cnbattle/hello-micro/pkg/jwt"

// RefreshToken 刷新token
func RefreshToken(token string) (string, error) {
	return jwt.RefreshToken(token)
}
