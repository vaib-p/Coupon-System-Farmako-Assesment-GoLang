package main

import (
	"coupon-system/config"
	"coupon-system/repository"
	"coupon-system/router"

	"github.com/gin-gonic/gin"
)

// @title Coupon System API
// @version 1.0
// @description API documentation for Coupon System
// @host localhost:8080
// @BasePath /api

func main() {

	config.LoadConfig()
	repository.InitDB()
	r := gin.Default()
	router.InitRoutes(r)

	r.Run(":" + config.AppConfig.Port)
}
