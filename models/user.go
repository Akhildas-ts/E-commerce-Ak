package models

type SignupDetail struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

type SignupDetailResponse struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type TokenUser struct {
	Users        SignupDetailResponse
	AccessToken  string
	RefreshToken string
}

type UserLogin struct {
	Email     string `json:"email" validate:"email"`
	Passoword string `json:"password" validate:"min=8,max=20"`
}

type UserLoginResponse struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type AddressInfo struct {
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
	Street    string `json:"street"`
	City      string `json:"city"`
}

type UsersProfileDetails struct {
	Firstname string `json:"firstname"  `
	Lastname  string `json:"lastname" `
	Email     string `json:"email" `
	Phone     string `json:"phone" `
	// ReferralCode string `json:"referral_code" binding:"required"`

}

type AddressInfoResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
	Street    string `json:"street"`
	City      string `json:"city"`
}

type PaymentDetails struct {
	ID           uint   `json:"id"`
	Payment_Name string `json:"payment_name"`
}
   
type CheckoutDetails struct {
	AddressInfoResponse []AddressInfoResponse
	Payment_Method      []PaymentDetails
	Cart                []Cart
	Grand_Total         float64
	Total_Price         float64
	DiscountReason      []string
}


type ForgetPassword struct { 

ReEnterPassword  string  `json:"re_password"`
ConformPassword string   `json:"conform"`

}


type UpdatePassword struct {
	OldPassword        string `json:"old_password" binding:"required"`
	NewPassword        string `json:"new_password" binding:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" binding:"required"`
}

// type BillingAddresInfo struct {
// 	Name      string `json:"name" validate:"required"`
// 	HouseName string `json:"house_name" validate:"required"`
// 	State     string `json:"state" validate:"required"`
// 	Pin       string `json:"pin" validate:"required"`
// 	Street    string `json:"street"`
// 	City      string `json:"city"`
// }
