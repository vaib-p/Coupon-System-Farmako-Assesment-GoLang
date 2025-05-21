package handlers

import (
	"coupon-system/cache"
	"coupon-system/models"
	"coupon-system/repository"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	gocache "github.com/patrickmn/go-cache"
)

// CreateCoupon godoc
// @Summary Create a new coupon
// @Description Create a coupon with necessary details
// @Tags coupons
// @Accept json
// @Produce json
// @Param coupon body models.Coupon true "Coupon Data"
// @Success 201 {object} models.CouponSwagger
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /coupons [post]
func CreateCoupon(c *gin.Context) {
	ctx := c.Request.Context()

	var coupon models.Coupon
	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request body"})
		return
	}

	if err := repository.DB.WithContext(ctx).Create(&coupon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create coupon"})
		return
	}

	c.JSON(http.StatusCreated, coupon)
}

// GetAllCoupons godoc
// @Summary Get all coupons
// @Description Retrieve a list of all coupons
// @Tags coupons
// @Produce json
// @Success 200 {array} models.CouponSwagger
// @Failure 500 {object} models.ErrorResponse
// @Router /coupons [get]
func GetAllCoupons(c *gin.Context) {
	ctx := c.Request.Context()

	var coupons []models.Coupon
	if err := repository.DB.WithContext(ctx).Find(&coupons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to retrieve coupons"})
		return
	}
	c.JSON(http.StatusOK, coupons)
}

// ValidateCoupon godoc
// @Summary Validate a coupon code
// @Description Checks if the coupon code is valid for the given total and time
// @Tags coupons
// @Accept json
// @Produce json
// @Param coupon body models.ValidateCouponRequest true "Coupon Validation Request"
// @Success 200 {object} models.CouponValidationResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /coupons/validate [post]
func ValidateCoupon(c *gin.Context) {
	ctx := c.Request.Context()

	var request struct {
		CouponCode string  `json:"coupon_code"`
		Total      float64 `json:"total"`
		Time       string  `json:"time"` // format: "15:04"
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request"})
		return
	}

	cacheKey := fmt.Sprintf("%s-%.2f-%s", request.CouponCode, request.Total, request.Time)

	if cached, found := cache.CouponCache.Get(cacheKey); found {
		c.JSON(http.StatusOK, cached)
		return
	}

	var coupon models.Coupon
	err := repository.DB.WithContext(ctx).First(&coupon, "coupon_code = ?", request.CouponCode).Error
	if err != nil {
		response := models.CouponValidationResponse{
			IsValid: false,
			Reason:  "Coupon not found",
		}
		cache.CouponCache.Set(cacheKey, response, gocache.DefaultExpiration)
		c.JSON(http.StatusOK, response)
		return
	}

	if coupon.ExpiryDate.Before(time.Now()) {
		response := models.CouponValidationResponse{
			IsValid: false,
			Reason:  "Coupon expired",
		}
		cache.CouponCache.Set(cacheKey, response, gocache.DefaultExpiration)
		c.JSON(http.StatusOK, response)
		return
	}

	if request.Total < coupon.MinOrderValue {
		response := models.CouponValidationResponse{
			IsValid: false,
			Reason:  "Minimum order value not met",
		}
		cache.CouponCache.Set(cacheKey, response, gocache.DefaultExpiration)
		c.JSON(http.StatusOK, response)
		return
	}

	layout := "15:04"
	start, _ := time.Parse(layout, coupon.ValidTimeWindow[0])
	end, _ := time.Parse(layout, coupon.ValidTimeWindow[1])
	current, _ := time.Parse(layout, request.Time)

	if current.Before(start) || current.After(end) {
		response := models.CouponValidationResponse{
			IsValid: false,
			Reason:  "Coupon not valid at this time",
		}
		cache.CouponCache.Set(cacheKey, response, gocache.DefaultExpiration)
		c.JSON(http.StatusOK, response)
		return
	}

	response := models.CouponValidationResponse{
		IsValid: true,
		Message: "Coupon is valid",
		Discount: &models.Discount{
			Type:  string(coupon.DiscountType),
			Value: coupon.DiscountValue,
		},
	}
	cache.CouponCache.Set(cacheKey, response, gocache.DefaultExpiration)
	c.JSON(http.StatusOK, response)
}
