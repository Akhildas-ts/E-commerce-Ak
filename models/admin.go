package models

type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=8,max=20"`
}

type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"  `
	Lastname  string `json:"lastname" `
	Email     string `json:"email" `
}


type DashBoardUser struct{
	TotalUsers int
    BlockedUser int
	UnBlockedUser int
}



type CompleteAdminDashBoard struct {
	DashBoarduser DashBoardUser
	
}