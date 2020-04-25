package bootstrap

import (
	"errors"
	"fmt"
	"dofun/app/models/category"
	"dofun/app/models/image"
	"dofun/app/models/link"
	"dofun/app/models/notification"
	passwordreset "dofun/app/models/password_reset"
	"dofun/app/models/permission"
	"dofun/app/models/reply"
	"dofun/app/models/topic"
	"dofun/app/models/user"
	"dofun/config"
	"dofun/database"
	"dofun/database/factory"

	"github.com/jinzhu/gorm"
)


// SetupDB db setup
func SetupDB(needMock bool) (*gorm.DB, error) {
	db := database.InitDB()

	// db migrate
	db.AutoMigrate(
		// permission
		&permission.Permission{},
		&permission.Role{},
		&permission.ModelHasPermission{},
		&permission.ModelHasRole{},
		&permission.RoleHasPermission{},

		&user.User{},
		&passwordreset.PasswordReset{},
		&category.Category{},
		&topic.Topic{},
		&reply.Reply{},
		&notification.Notification{},
		&image.Image{},
		&link.Link{},
	)

	// 只有非 release 时才可 mock
	if needMock {
		if config.AppConfig.RunMode == config.RunmodeRelease {
			panic("[mock] 请在非生产环境中使用")
		}

		fmt.Print("\n\n-------------------------------------------------- MOCK --------------------------------------------------\n\n")
		factory.Mock()
		return db, errors.New("mock data")
	}

	return db, nil
}
