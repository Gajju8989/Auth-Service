package wire

import (
	"github.com/google/wire"
	"github/com/Gajju8989/Auth_Service/internal/handler"
	"github/com/Gajju8989/Auth_Service/internal/repo"
	"github/com/Gajju8989/Auth_Service/internal/router"
	"github/com/Gajju8989/Auth_Service/internal/service"
	"gorm.io/gorm"
)

func InitializeAuthServiceNew(db *gorm.DB) (router.MapRouter, error) {
	wire.Build(repo.NewRepository, service.NewAuthService, handler.NewHandler, router.NewRouter)
	return nil, nil
}
