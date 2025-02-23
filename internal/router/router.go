package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/com/Gajju8989/Auth_Service/internal/handler"
	"github/com/Gajju8989/Auth_Service/internal/middleware"
)

type MapRouter interface {
	SetupRoutes(engine *gin.Engine)
}

type router struct {
	handler handler.Handler
}

func NewRouter(handler handler.Handler) MapRouter {
	return &router{handler: handler}
}

func (r *router) SetupRoutes(engine *gin.Engine) {
	api := engine.Group("/api")
	v1 := api.Group("/v1")
	auth := v1.Group("/auth")
	{
		auth.POST("/signup", r.handler.CreateUser)
		auth.POST("/login", r.handler.Login)
		auth.POST("/refresh-token", r.handler.RefreshToken)
		auth.POST("/revoke-token", r.handler.RevokeToken)
	}

	protected := v1.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", r.handler.GetProfile)
	}

	err := engine.Run(":8080")
	if err != nil {
		fmt.Print("Unable To Run The Server")
		return
	}
}
