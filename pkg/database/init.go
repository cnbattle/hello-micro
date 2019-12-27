package database

import (
	"fmt"
	"github.com/cnbattle/hello-micro/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strings"
)

var (
	Conn *gorm.DB
)

func init() {
	var err error
	var dsn string
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s",
		config.GetDefaultEnv("DB_USERNAME", "root"),
		config.GetDefaultEnv("DB_PASSWORD", ""),
		config.GetDefaultEnv("DB_HOST", "127.0.0.1"),
		config.GetDefaultEnv("DB_PORT", "3306"),
		config.GetDefaultEnv("DB_DATABASE", "root"),
		strings.Replace(config.GetDefaultEnv("APP_TIMEZONE", "Local"), "/", "%2F", -1),
	)
	Conn, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Panicf("init DB err:%v,dsn:%v", err, dsn)
	}
	// 全局禁用表名复数
	Conn.SingularTable(true)
	//启用Logger，显示详细日志，默认情况下会打印发生的错误
	//Conn.LogMode(false)
	Conn.DB().SetMaxIdleConns(10)
	Conn.DB().SetMaxOpenConns(100)
}
