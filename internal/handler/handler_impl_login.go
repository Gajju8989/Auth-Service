package handler

import (
	"github.com/gin-gonic/gin"
	"github/com/Gajju8989/Auth_Service/internal/service"
	"github/com/Gajju8989/Auth_Service/internal/service/model"
	"net/http"
)

func (h *impl) Login(ctx *gin.Context) {
	var req *model.UserAuthRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{errorKey: err.Error()})
		return
	}

	resp, err := h.service.Login(ctx, req)
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

	ctx.JSON(http.StatusCreated, resp)
}
