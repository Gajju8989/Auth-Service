package migration

import (
	"github/com/Gajju8989/Auth_Service/internal/repo/model/refreshtoken"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/token"
	"github/com/Gajju8989/Auth_Service/internal/repo/model/user"
	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) error {
	if err := db.AutoMigrate(&user.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&token.AccessToken{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&refreshtoken.RefreshToken{}); err != nil {
		return err
	}

	return nil
}
