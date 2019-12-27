package database

import (
	"github.com/cnbattle/hello-micro/pkg/database"
	"github.com/cnbattle/hello-micro/srv/user-srv/models"
	"github.com/jinzhu/gorm"
)

var UserConn *gorm.DB

func init() {
	UserConn = database.Conn
	UserConn.AutoMigrate(&models.UserAuth{})
	UserConn.AutoMigrate(&models.UserBase{})
}
