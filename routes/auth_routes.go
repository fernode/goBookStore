package routes

import (
	"github.com/fernode/goBookStore/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/register", controllers.RegisterUserController)

	router.POST("/login", controllers.Login)
}
