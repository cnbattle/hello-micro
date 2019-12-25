package jwt

import (
	"github.com/cnbattle/hello-micro/kernel/utils"
	"github.com/cnbattle/hello-micro/srv/user-srv/models"
	"github.com/google/uuid"
	"math/rand"
	"testing"
)

// TestGenerateToken 生成方法
func TestGenerateToken(t *testing.T) {
	type args struct {
		base *models.UserBase
	}
	type testsStruct struct {
		name string
		args args
	}
	var tests []testsStruct

	for i := 1; i < 10; i++ {
		var userBase models.UserBase
		userBase.Uid = uuid.New().String()
		userBase.Nickname = utils.RandString(6)
		userBase.Avatar = "https://ruan.co/" + utils.RandString(6)
		userBase.Gender = rand.Intn(2)
		userBase.Signature = utils.RandString(128)
		tests = append(tests, testsStruct{
			name: "randData",
			args: args{&userBase},
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateToken(tt.args.base)
			if err != nil {
				t.Errorf("GenerateToken() error = %v", err)
				return
			}
		})
	}
}

// TestParseToken 解析token
func TestParseToken(t *testing.T) {
	oldToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzcyNjM2ODYsImlhdCI6MTU3NzI2MzY3NiwidWlkIjoidWlkIiwibmlja25hbWUiOiJuaWNrbmFtZSIsImF2YXRhciI6ImF2YXRhciIsImdlbmRlciI6ImdlbmRlciIsInNpZ25hdHVyZSI6InNpZ25hdHVyZSIsInBlcm1pc3Npb25zIjpudWxsfQ.jUXsL4WMM2lJKu56DMHpHTwtPGvAixQwfxxlyI9X0bM"
	_, err := ParseToken(oldToken)
	if err == nil {
		t.Error("testing oldToken is not error")
		return
	}
	var userBase models.UserBase
	userBase.Uid = uuid.New().String()
	userBase.Nickname = utils.RandString(6)
	userBase.Avatar = "https://ruan.co/" + utils.RandString(6)
	userBase.Gender = rand.Intn(2)
	userBase.Signature = utils.RandString(128)
	generateToken, err := GenerateToken(&userBase)
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
		var userBase models.UserBase
		userBase.Uid = uuid.New().String()
		userBase.Nickname = utils.RandString(6)
		userBase.Avatar = "https://ruan.co/" + utils.RandString(6)
		userBase.Gender = rand.Intn(2)
		userBase.Signature = utils.RandString(128)
		generateToken, err := GenerateToken(&userBase)
		if err != nil {
			t.Errorf("GenerateToken() error: %v", err)
		}
		tests = append(tests, testsStruct{
			name: "randData",
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
