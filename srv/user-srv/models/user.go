package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserAuth struct {
	gorm.Model
	UID          string `gorm:"type:char(36);not null;unique_index"` // 用户ID
	IdentityType string `gorm:"type:char(16);not null"`              // 身份标识ID
	Identity     string `gorm:"type:char(32);not null"`              // 身份标识
	Certificate  string `gorm:"type:char(32);not null"`              // 凭证
}

type UserBase struct {
	UID       string `gorm:"type:char(36);not null;unique_index"` // 用户ID
	UserRole  int    `gorm:"type:tinyint(2);not null"`            // 用户角色ID
	Nickname  string `gorm:"type:varchar(16);not null;unique"`    // 用户昵称
	Avatar    string `gorm:"type:varchar(128);not null"`          // 用户头像
	Gender    int    `gorm:"type:tinyint(1);default:0;not null"`  // 用户性别
	Signature string `gorm:"type:varchar(128);not null"`          // 用户签名
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
