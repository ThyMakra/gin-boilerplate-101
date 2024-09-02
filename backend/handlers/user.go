package handlers

import (
	"net/http"

	"github.com/ThyMakra/gin-boilerplate/backend/schemas"
	"github.com/ThyMakra/gin-boilerplate/backend/services"
	"github.com/ThyMakra/gin-boilerplate/pkg/utils"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	user services.UserEntity
}

func (h *userHandler) PingHandler(ctx *gin.Context) {
	utils.ApiResponse(ctx, "User route is reachable", http.StatusOK, nil)
}

func NewUserHandler(user services.UserEntity) *userHandler {
	return &userHandler{user: user}
}

func (h *userHandler) RegisterHandler(ctx *gin.Context) {
	var body schemas.UserSchema
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		utils.ApiResponse(ctx, "Parse json from body failed", http.StatusBadRequest, nil)
		return
	}

	_, error := h.user.RegisterEntity(&body)

	if error.Code != 0 {
		utils.ApiResponse(ctx, error.Message, error.Code, nil)
		return
	}

	utils.ApiResponse(ctx, "Register new user account success", http.StatusOK, nil)
}

func (h *userHandler) LoginHanlder(ctx *gin.Context) {
	var body schemas.UserSchema
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		utils.ApiResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	res, error := h.user.LoginEntity(&body)

	if error.Code != 0 {
		utils.ApiResponse(ctx, error.Message, error.Code, nil)
		return

	}

	utils.ApiResponse(ctx, "Login successfully", http.StatusOK, gin.H{
		"accessToken": nil,
		"expiredAt":   nil,
		"userId":      res.ID,
	})
}
