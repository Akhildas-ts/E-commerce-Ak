package models

// type Product struct {
// 	ID                uint     `json:"id"   gorm:"unique; not null"`
// 	Name              string   `json:"name"   gorm:"unique; not null" validate:"required"`
// 	SKU               string   `json:"sku"`
// 	CategoryID        uint     `json:"category_id"`
// 	Category          Category `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
// 	DesignDescription string   `json:"design_description"`
// 	BrandID           uint     `json:"brand_id"`
// 	Quantity          int      `json:"quantity"`
// 	Price             float64  `json:"price"`
// 	ProductStatus     string   `json:"product_status"`
// }

// type Category struct {
// 	ID           uint   `json:"id" gorm:"unique; not null"`
// 	CategoryName string `json:"category_name"  gorm:"unique; not null"  validate:"required"`
// }

type ProductReceiver struct {
	Name              string  `json:"name"`
	SKU               string  `json:"sku"`
	CategoryId        string  `json:"category_id"`
	DesignDescription string  `json:"design_description"`
	BrandID           uint    `json:"brand_id"`
	Quantity          int     `json:"quantity"`
	Price             float64 `json:"price"`
	ProductStatus     string  `json:"product_status"`
}

type ProductUpdate struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
type ProductUpdateReciever struct {
	ProductID int
	Quantity  int
}
