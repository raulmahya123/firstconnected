package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func NilaiRoute(router *gin.Engine) {
	router.POST("/nilai", controllers.CreateNilai())
	router.GET("/nilai/:nialiGetID", controllers.GetNilai())
	router.PUT("/nilai/:nialiID", controllers.EditNilai())
	router.DELETE("/nilai/:nialiID", controllers.DeleteNilai())
	router.GET("/nilais", controllers.GetAllNilais())
	router.GET("/nilaii/:npm", controllers.GetNPM())
}
