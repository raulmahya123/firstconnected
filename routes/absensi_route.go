package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func AbsensiRoute(router *gin.Engine) {

	router.POST("/asbensi", controllers.CreateAbsensi())
	router.GET("/absensi/:absensiGetId", controllers.GetAabsensi())
	router.PUT("/absensi/:absensiID", controllers.EditAbsensi())
	router.DELETE("/absensi/:absensiID", controllers.DeleteAabsensi())
	router.GET("/absensis", controllers.GetAllAbssenis())
}
