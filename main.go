package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes" //add this
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//routes
	routes.AbsensiRoute(router)   //add this
	routes.OrangTuaRoute(router)  //add this
	routes.NilaiRoute(router)     //add this
	routes.MahasiswaRoute(router) //add this
	routes.MataKuliahRoute(router)

	router.Run("localhost:8080")
}
