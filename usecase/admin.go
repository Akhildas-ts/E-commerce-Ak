package usecase

import (
	"ak/domain"
	"ak/helper"
	"ak/models"
	"ak/repository"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(adminmodel models.AdminLogin) (domain.TokenAdmin, error) {

	// GETTING DETAILS FROM ADMIN WITH EMAIL

	AdminDetail, err := repository.AdminLogin(adminmodel)
	if err != nil {
		fmt.Println("error:1")
		return domain.TokenAdmin{}, errors.New("demo : repository adminlogin")
		
	}

    password := AdminDetail.Password

	fmt.Println(password)
	fmt.Println("Stored Hashed Password:", AdminDetail.Password)
	err = bcrypt.CompareHashAndPassword([]byte(AdminDetail.Password), []byte(adminmodel.Password))
	if err != nil {
		fmt.Println("error:2")
		return domain.TokenAdmin{}, errors.New("demo : repository adminlogin")

	}

	var AdminDetailsResponse models.AdminDetailsResponse

	tokenString, err := helper.GenerateTokenAdmin(AdminDetailsResponse)

	if err != nil {
		fmt.Println("error:3")
		return domain.TokenAdmin{},  errors.New("demo : repository adminlogin")

	}

	return domain.TokenAdmin{
		Admin: AdminDetailsResponse,
		Token: tokenString,
	}, nil

}

func DashBoard() (models.CompleteAdminDashBoard, error) {

	userDetails, err := repository.DashBoardUserDetails()

	if err != nil {
		return models.CompleteAdminDashBoard{}, nil
	}

	return models.CompleteAdminDashBoard{
		DashBoarduser: userDetails,
	}, nil

}
