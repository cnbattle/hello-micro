package loginservice

import (
	"errors"
	"github.com/cnbattle/hello-micro/pkg/database"
	"github.com/cnbattle/hello-micro/srv/user-srv/models"

	"github.com/google/uuid"
	"github.com/silenceper/wechat/util"
)

// LoginOrCreate
func loginOrCreate(identityType, identity string, autoCreate bool) (*models.UserAuth, error) {
	userAuth := models.UserAuth{}
	err := database.Conn.Where("identity_type=?", identityType).
		Where("identity=?", identity).First(&userAuth).Error
	if err != nil {
		if autoCreate {
			userAuth.IdentityType = identityType
			userAuth.Identity = identity
			userAuth.UID = uuid.New().String()
			database.Conn.Create(&userAuth)
		} else {
			return &models.UserAuth{}, errors.New("not has user auth data")
		}
	}
	userAuth.Certificate = util.MD5Sum(uuid.New().String())
	database.Conn.Save(&userAuth)
	return &userAuth, nil
}
