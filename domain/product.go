package domain

import "gorm.io/gorm"

type Products struct {
	//validator use
	*gorm.Model       `json:"-"`
	ID                uint     `json:"id"   gorm:"unique; not null"`
	Name              string   `json:"name"   gorm:"unique; not null" validate:"required"`
	SKU               string   `json:"sku"`
	CategoryID        uint     `json:"category_id"`
	Category          Category `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	DesignDescription string   `json:"design_description"`
	BrandID           uint     `json:"brand_id"`
	Quantity          int      `json:"quantity"`
	Price             float64  `json:"price"`
	ProductStatus     string   `json:"product_status"`
	IsDeleted         bool     `json:"is_deleted" gorm:"default:false"`
}

type Category struct {
	ID           uint   `json:"id" gorm:"unique; not null"`
	CategoryName string `json:"category_name"  gorm:"unique; not null"`
}
