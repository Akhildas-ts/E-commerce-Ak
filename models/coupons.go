package models 


type AddCoupon struct {
	Coupon             string  `json:"coupon" binding:"required" validate:"required"`
	DiscountPercentage int     `json:"discount_percentage" binding:"required"`
	MinimumPrice       float64 `json:"minimum_price" binding:"required"`
	Validity           bool    `json:"validity" binding:"required"`
}

type Coupon struct {
	ID                 uint    `json:"id"`
	Coupon             string  `json:"coupon"`
	DiscountPercentage int     `json:"discount_percentage"`
	MinimumPrice       float64 `json:"minimum_price"`
	Validity           bool    `json:"validity"`
}

type CouponAddUser struct {
	CouponName string `json:"coupon_name" binding:"required"`
}

type ProductOfferBriefResponse struct {
	ProductsBrief ProductBrief
	OfferResponse OfferResponse
}
type OfferResponse struct {
	OfferID         uint    `json:"offer_id"`
	OfferName       string  `json:"offer_name"`
	OfferPercentage int     `json:"offer_percentage"`
	OfferPrice      float64 `json:"offer_price"`
	OfferType       string  `json:"offer_type"`
	OfferLimit      int     `json:"offer_limit"`
}

type ProductOfferReceiver struct {
	ProductID          uint   `json:"product_id" binding:"required"`
	OfferName          string `json:"offer_name" binding:"required"`
	DiscountPercentage int    `json:"discount_percentage" binding:"required"`
	OfferLimit         int    `json:"offer_limit"`
}

type CategoryOfferReceiver struct {
	CategoryID         uint   `json:"category_id" `
	OfferName          string `json:"offer_name"`
	DiscountPercentage int    `json:"discount_percentage"`
	OfferLimit         int    `json:"offer_limit"`
}