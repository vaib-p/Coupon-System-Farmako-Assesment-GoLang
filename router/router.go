package router

import (
	"coupon-system/handlers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "coupon-system/docs" // swag generated docs

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := r.Group("/api")
	{
		api.POST("/coupons", handlers.CreateCoupon)
		api.GET("/coupons", handlers.GetAllCoupons)
		api.POST("/coupons/validate", handlers.ValidateCoupon)
	}

}
