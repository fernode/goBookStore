package routes

import (
	"github.com/fernode/goBookStore/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUserController)
	router.POST("/login", controllers.Login)

	router.GET("/user", controllers.GetUserProfile)
	router.PUT("/user", controllers.UpdateUserProfile)
	router.DELETE("/user", controllers.DeleteUserProfile)
}
