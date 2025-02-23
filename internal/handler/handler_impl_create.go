package handler

import (
	"github.com/gin-gonic/gin"
	"github/com/Gajju8989/Auth_Service/internal/service"
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

	ctx.JSON(http.StatusCreated, resp)
}
