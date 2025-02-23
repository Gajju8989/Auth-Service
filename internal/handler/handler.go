package handler

import (
	"github.com/gin-gonic/gin"
	"github/com/Gajju8989/Auth_Service/internal/service"
)

type Handler interface {
	CreateUser(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
	RevokeToken(ctx *gin.Context)
}

type impl struct {
	service service.AuthService
}

func NewHandler(service service.AuthService) Handler {
	return &impl{
		service: service,
	}
}
