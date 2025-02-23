package main

import (
	"github.com/gin-gonic/gin"
	config "github/com/Gajju8989/Auth_Service/internal/config/database"
	"github/com/Gajju8989/Auth_Service/internal/config/database/migration"
	handler2 "github/com/Gajju8989/Auth_Service/internal/handler"
	repo2 "github/com/Gajju8989/Auth_Service/internal/repo"
	router2 "github/com/Gajju8989/Auth_Service/internal/router"
	service2 "github/com/Gajju8989/Auth_Service/internal/service"
)

func main() {
	config.InitDB()
	err := migration.MigrateAll(config.GetDB())
	if err != nil {
		return
	}

	repo := repo2.NewRepository(config.GetDB())
	service := service2.NewAuthService(repo)
	handler := handler2.NewHandler(service)
	router := router2.NewRouter(handler)
	r := gin.Default()
	router.SetupRoutes(r)
}
