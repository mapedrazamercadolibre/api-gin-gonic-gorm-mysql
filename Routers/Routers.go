package Routers

import (
	"api-gin-gionic-gorm-mysql/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	group := router.Group("/user-api")
	{
		group.GET("users", Controllers.GetUsers)
		group.POST("users", Controllers.CreateUser)
		group.GET("users/:id", Controllers.GetUserByID)
		group.PUT("users/:id", Controllers.UpdateUser)
		group.DELETE("users/:id", Controllers.DeleteUser)
	}
	return router
}
