package routes

import (
	"github.com/ThyMakra/gin-boilerplate/backend/handlers"
	"github.com/ThyMakra/gin-boilerplate/backend/repositories"
	"github.com/ThyMakra/gin-boilerplate/backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRoute(db *gorm.DB, router *gin.Engine) {
	userRepository := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(services.NewUserService(userRepository))

	userRoute := router.Group("/api/v1/auth")

	userRoute.GET("/ping", userHandler.PingHandler)
	userRoute.POST("/register", userHandler.RegisterHandler)
	userRoute.POST("/login", userHandler.LoginHanlder)
}
