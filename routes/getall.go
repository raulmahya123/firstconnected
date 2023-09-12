package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func GetallRoute(router *gin.Engine) {
	router.GET("/getall", controllers.GetAllDB())
}
