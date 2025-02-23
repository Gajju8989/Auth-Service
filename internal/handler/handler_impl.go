package handler

import (
	"github.com/gin-gonic/gin"
	"github/com/Gajju8989/Auth_Service/internal/middleware"
	"github/com/Gajju8989/Auth_Service/internal/service/model"
	"net/http"
)

func (h *impl) CreateUser(ctx *gin.Context) {
	var req *model.UserAuthRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *impl) Login(ctx *gin.Context) {
	var req *model.UserAuthRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.Login(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *impl) GetProfile(ctx *gin.Context) {
	resp, err := h.service.GetProfiles(ctx, middleware.GetUserID(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *impl) RefreshToken(ctx *gin.Context) {
	var req *model.RefreshTokenRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.RefreshToken(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *impl) RevokeToken(ctx *gin.Context) {
	var req *model.RevokeTokenRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.RevokeToken(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"message": "Token revoked successfully"})
}
