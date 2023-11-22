package repository

import (
	"ak/database"
	"ak/models"
)

func UserDetailsUsingPhone(phone string) (models.SignupDetailResponse,error) {

	var userDetails models.SignupDetailResponse

	if err := database.DB.Raw("select * from users where phone = ?",phone).Scan(&userDetails).Error;err != nil {
	   return models.SignupDetailResponse{},err
	}
	return userDetails,nil
}