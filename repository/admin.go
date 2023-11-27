package repository

import (
	"ak/database"
	"ak/domain"
	"ak/models"
	"errors"
	"fmt"
)

func AdminLogin(admin models.AdminLogin) (domain.Admin, error) {
	var admindomain domain.Admin

	if err := database.DB.Raw("select * from users where email= ? and isadmin= true ", admin.Email).Scan(&admindomain).Error; err != nil {
		return domain.Admin{}, errors.New("admin email is not available on database")
	}
	fmt.Println(admindomain.Password)

	return admindomain, nil
}

func DashBoardUserDetails() (models.DashBoardUser, error) {

	var userDetails models.DashBoardUser

	err := database.DB.Raw("select count (*) from users").Scan(&userDetails.TotalUsers).Error

	if err != nil {

		return models.DashBoardUser{}, nil

	}

	err = database.DB.Raw("select count(*) from users where blocked = true").Scan(&userDetails.BlockedUser).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	fmt.Println("hi")

	return userDetails, nil
}
