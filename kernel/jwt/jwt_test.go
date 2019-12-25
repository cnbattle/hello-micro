package jwt

import (
	"github.com/cnbattle/hello-micro/kernel/utils"
	"github.com/google/uuid"
	"log"
	"testing"
)

// TestGenerateToken
func TestGenerateToken(t *testing.T) {
	uid := uuid.New().String()
	nickname := utils.RandString(6)
	avatar := "https://ruan.co/" + utils.RandString(6)
	gender := utils.RandString(1)
	signature := utils.RandString(128)
	generateToken, err := GenerateToken(uid, nickname, avatar, gender, signature)
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
	uid := uuid.New().String()
	nickname := utils.RandString(6)
	avatar := "https://ruan.co/" + utils.RandString(6)
	gender := utils.RandString(1)
	signature := utils.RandString(128)
	generateToken, err := GenerateToken(uid, nickname, avatar, gender, signature)
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
	type testsStruct struct {
		name string
		args string
		want bool
	}
	var tests []testsStruct

	tests = append(tests, testsStruct{
		name: "error",
		args: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzcyNjM2ODYsImlhdCI6MTU3NzI2MzY3NiwidWlkIjoidWlkIiwibmlja25hbWUiOiJuaWNrbmFtZSIsImF2YXRhciI6ImF2YXRhciIsImdlbmRlciI6ImdlbmRlciIsInNpZ25hdHVyZSI6InNpZ25hdHVyZSIsInBlcm1pc3Npb25zIjpudWxsfQ.jUXsL4WMM2lJKu56DMHpHTwtPGvAixQwfxxlyI9X0bM",
		want: false,
	})
	for i := 0; i < 10; i++ {
		uid := uuid.New().String()
		nickname := utils.RandString(6)
		avatar := "https://ruan.co/" + utils.RandString(6)
		gender := utils.RandString(1)
		signature := utils.RandString(128)
		generateToken, err := GenerateToken(uid, nickname, avatar, gender, signature)
		if err != nil {
			t.Errorf("GenerateToken() error: %v", err)
		}
		tests = append(tests, testsStruct{
			name: "true",
			args: generateToken,
			want: true,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyToken(tt.args); got != tt.want {
				t.Errorf("VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
