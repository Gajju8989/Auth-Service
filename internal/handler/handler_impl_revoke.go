package handler

import (
	"github.com/gin-gonic/gin"
	"github/com/Gajju8989/Auth_Service/internal/middleware"
	"github/com/Gajju8989/Auth_Service/internal/service"
	"net/http"
)

func (h *impl) RevokeToken(ctx *gin.Context) {
	err := h.service.RevokeToken(ctx, middleware.GetUserID(ctx))
	if err != nil {
		genericErr, ok := err.(*service.GenericResponse)
		if ok {
			ctx.JSON(genericErr.StatusCode, genericErr)
		} else {
			ctx.JSON(http.StatusInternalServerError, service.GenericResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    internalServerErrorMessage,
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, gin.H{messageKey: tokenRevokedSuccessfully})
}
