package usecase

import (
	"ak/domain"
	"ak/helper"
	"ak/models"
	"ak/repository"

	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(adminmodel models.AdminLogin) (domain.TokenAdmin, error) {

	// GETTING DETAILS FROM ADMIN WITH EMAIL

	AdminDetail, err := repository.AdminLogin(adminmodel)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(AdminDetail.Password), []byte(adminmodel.Password)) 
	if err != nil {
		return domain.TokenAdmin{},err

	
	}

	var AdminDetailsResponse models.AdminDetailsResponse

	tokenString,err := helper.GenerateTokenAdmin(AdminDetailsResponse) 

	if err != nil {
		return domain.TokenAdmin{},err

	}

	return domain.TokenAdmin{
		Admin: AdminDetailsResponse,
		Token: tokenString,
	},nil



}


func DashBoard()(models.CompleteAdminDashBoard,error){
  

	userDetails,err := repository.DashBoardUserDetails()

	if err != nil {
		return models.CompleteAdminDashBoard{},nil
	}

	return models.CompleteAdminDashBoard{
		DashBoarduser: userDetails,
	},nil
	
}
