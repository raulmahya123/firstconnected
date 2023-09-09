package routes

import (
	"gin-mongo-api/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AbsensiRoute(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.POST("/asbensi", controllers.CreateAbsensi())
	router.GET("/absensi/:absensiGetId", controllers.GetAabsensi())
	router.PUT("/absensi/:absensiID", controllers.EditAbsensi())
	router.DELETE("/absensi/:absensiID", controllers.DeleteAabsensi())
	router.GET("/absensis", controllers.GetAllAbssenis())
}
