package models

import (
	"time"

	"github.com/lib/pq"
)

type UsageType string

const (
	OneTime   UsageType = "one_time"
	MultiUse  UsageType = "multi_use"
	TimeBased UsageType = "time_based"
)

type DiscountType string

const (
	Flat       DiscountType = "flat"
	Percentage DiscountType = "percentage"
)

type Coupon struct {
	CouponCode            string         `json:"coupon_code" gorm:"primaryKey"`
	ExpiryDate            time.Time      `json:"expiry_date"`
	UsageType             UsageType      `json:"usage_type"`
	ApplicableMedicineIDs pq.StringArray `json:"applicable_medicine_ids" gorm:"type:text[]"`
	ApplicableCategories  pq.StringArray `json:"applicable_categories" gorm:"type:text[]"`
	MinOrderValue         float64        `json:"min_order_value"`
	ValidTimeWindow       pq.StringArray `json:"valid_time_window" gorm:"type:text[]"`
	TermsAndConditions    string         `json:"terms_and_conditions"`
	DiscountType          DiscountType   `json:"discount_type"`
	DiscountValue         float64        `json:"discount_value"`
	MaxUsagePerUser       int            `json:"max_usage_per_user"`
}
type CouponSwagger struct {
	CouponCode            string    `json:"coupon_code" example:"SAVE20"`
	ExpiryDate            time.Time `json:"expiry_date" example:"2025-12-31T23:59:59Z"`
	UsageType             string    `json:"usage_type" example:"single_use"`
	ApplicableMedicineIDs []string  `json:"applicable_medicine_ids" example:"[\"med1\", \"med2\"]"`
	ApplicableCategories  []string  `json:"applicable_categories" example:"[\"category1\", \"category2\"]"`
	MinOrderValue         float64   `json:"min_order_value" example:"50.0"`
	ValidTimeWindow       []string  `json:"valid_time_window" example:"[\"09:00\", \"21:00\"]"`
	TermsAndConditions    string    `json:"terms_and_conditions" example:"Use this coupon once per user."`
	DiscountType          string    `json:"discount_type" example:"percentage"`
	DiscountValue         float64   `json:"discount_value" example:"20.0"`
	MaxUsagePerUser       int       `json:"max_usage_per_user" example:"1"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Discount struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type CouponValidationResponse struct {
	IsValid  bool      `json:"is_valid"`
	Reason   string    `json:"reason,omitempty"`
	Message  string    `json:"message,omitempty"`
	Discount *Discount `json:"discount,omitempty"`
}
type ValidateCouponRequest struct {
	CouponCode string  `json:"coupon_code" example:"SAVE20"`
	Total      float64 `json:"total" example:"150.75"`
	Time       string  `json:"time" example:"14:30"` // format HH:mm
}
