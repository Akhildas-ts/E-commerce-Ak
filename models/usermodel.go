package models

type SignupDetail struct {
	FirstName string `json:"firstname "`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

type SignupDetailResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname "`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
}

type TokenUser struct {
	Users        SignupDetailResponse
	AccessToken  string
	RefreshToken string
}

type UserLogin struct {
	Email string `json:"email" validate:"email"`
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

