package handler

import (
	"github.com/gin-gonic/gin"
	"github/com/Gajju8989/Auth_Service/internal/middleware"
	"github/com/Gajju8989/Auth_Service/internal/service"
	"net/http"
)

func (h *impl) RefreshToken(ctx *gin.Context) {
	resp, err := h.service.RefreshToken(ctx, middleware.GetUserID(ctx))
	if err != nil {
		genericErr, ok := err.(*service.GenericResponse)
		if ok {
			ctx.JSON(genericErr.StatusCode, genericErr)
		} else {
			ctx.JSON(http.StatusInternalServerError, service.GenericResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Internal server error",
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, resp)
}
