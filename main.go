package main

import (
	"github.com/gin-gonic/gin"
	config "github/com/Gajju8989/Auth_Service/internal/config/database"
	"github/com/Gajju8989/Auth_Service/internal/config/database/migration"
	"github/com/Gajju8989/Auth_Service/internal/wire"
	"log"
)

func main() {
	config.InitDB()
	err := migration.MigrateAll(config.GetDB())
	if err != nil {
		return
	}

	router, err := wire.InitializeAuthService(config.GetDB())
	if err != nil {
		log.Fatalf("failed to initialize router: %v", err)
	}

	engine := gin.Default()
	router.SetupRoutes(engine)
	err = engine.Run(":8080")
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
