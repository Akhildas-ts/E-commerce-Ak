package domain

import (
	"time"

	"gorm.io/gorm"
)

type Coupons struct {
	ID                 uint    `json:"id" gorm:"uniquekey; not null"`
	Coupon             string  `json:"coupon" gorm:"coupon"`
	DiscountPercentage int     `json:"discount_percentage"`
	Validity           bool    `json:"validity"`
	MinimumPrice       float64 `json:"minimum_price"`
}

type UsedCoupon struct {
	ID       uint    `json:"id" gorm:"uniquekey not null"`
	CouponID uint    `json:"coupon_id"`
	Coupons  Coupons `json:"-" gorm:"foreignkey:CouponID"`
	UserID   uint    `json:"user_id"`
	Users    User    `json:"-" gorm:"foreignkey:UserID"`
	Used     bool    `json:"used"`
}
type CategoryOffer struct {
	ID                 uint      `json:"id" gorm:"unique; not null"`
	CategoryID         uint      `json:"category_id"`
	Category           Category  `json:"-" gorm:"foreignkey:CategoryID"`
	OfferName          string    `json:"offer_name"`
	DiscountPercentage int       `json:"discount_percentage"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	OfferLimit         int       `json:"offer_limit"`
	OfferUsed          int       `json:"offer_used"`
}
type Referral struct {
	gorm.Model
	UserID         uint    `json:"user_id" gorm:"uniquekey; not null"`
	Users          User    `json:"-" gorm:"foreignkey:UserID"`
	ReferralCode   string  `json:"referral_code"`
	ReferralAmount float64 `json:"referral_amount"`
	ReferredUserID uint    `json:"referred_user_id"`
}
