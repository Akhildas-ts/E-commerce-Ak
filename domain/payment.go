package domain 



type PaymentMethod struct {
	ID           uint   `gorm:"primarykey"`
	Payment_Name string `json:"payment_name"`
}

type RazerPay struct {
	ID        uint   `json:"id" gorm:"primarykey not null"`
	OrderID   string `json:"order_id"`
	Order     Order  `json:"-" gorm:"foreignkey:OrderID"`
	RazorID   string `json:"razor_id"`
	PaymentID string `json:"payment_id"`
}
