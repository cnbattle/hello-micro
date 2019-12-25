package jwt

import (
	"log"
	"testing"
)

// TestGenerateToken
func TestGenerateToken(t *testing.T) {
	generateToken, err := GenerateToken("uid", "nickname", "avatar", "gender", "signature")
	if err != nil {
		t.Error(err)
	}
	log.Println(generateToken)
}

// TestParseToken 解析验证token
func TestParseToken(t *testing.T) {
	oldToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzcyNjM2ODYsImlhdCI6MTU3NzI2MzY3NiwidWlkIjoidWlkIiwibmlja25hbWUiOiJuaWNrbmFtZSIsImF2YXRhciI6ImF2YXRhciIsImdlbmRlciI6ImdlbmRlciIsInNpZ25hdHVyZSI6InNpZ25hdHVyZSIsInBlcm1pc3Npb25zIjpudWxsfQ.jUXsL4WMM2lJKu56DMHpHTwtPGvAixQwfxxlyI9X0bM"
	_, err := ParseToken(oldToken)
	if err == nil {
		t.Error("testing oldToken is not error")
		return
	}
	generateToken, err := GenerateToken("uid", "nickname", "avatar", "gender", "signature")
	if err != nil {
		t.Error(err)
	}
	_, err = ParseToken(generateToken)
	if err != nil {
		t.Error(err)
	}
}

// TestVerifyToken
func TestVerifyToken(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		// TODO: Add test cases.
		{"1", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzcyNjM2ODYsImlhdCI6MTU3NzI2MzY3NiwidWlkIjoidWlkIiwibmlja25hbWUiOiJuaWNrbmFtZSIsImF2YXRhciI6ImF2YXRhciIsImdlbmRlciI6ImdlbmRlciIsInNpZ25hdHVyZSI6InNpZ25hdHVyZSIsInBlcm1pc3Npb25zIjpudWxsfQ.jUXsL4WMM2lJKu56DMHpHTwtPGvAixQwfxxlyI9X0bM", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyToken(tt.args); got != tt.want {
				t.Errorf("VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
