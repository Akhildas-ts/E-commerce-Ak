package usecase

import (
	"ak/helper"
	"ak/models"
	"ak/repository"
	"errors"
)

func UsersingUp(user models.SignupDetail) (*models.TokenUser, error) {

	email, err := repository.CheckingEmailValidation(user.Email)

	if err != nil {
		return &models.TokenUser{}, errors.New("error with the singup server")

	}

	if email != nil {
		return &models.TokenUser{}, errors.New("email is already exisit ")
	}

	phone, err := repository.CheckingPhoneExists(user.Phone)

	if err != nil {

		return &models.TokenUser{}, errors.New("p_server have issue ")

	}
	if phone != nil {
		return &models.TokenUser{}, errors.New("phone number is already exist ")

	}

	//   Passoword Hash

	hashedPassword, err := helper.PasswordHashing(user.Password)

	if err != nil {
		return &models.TokenUser{}, errors.New("hash_server have issue")
	}

	user.Password = hashedPassword

	dataInsert, err := repository.SignupInsert(user)
	if err != nil {
		return &models.TokenUser{}, errors.New("could not add User ")
	}

	// CREATING A JWT TOKEN FOR THE NEW USER\\

	accessToken, err := helper.GenerateAccessToken(dataInsert)

	if err != nil {
		return &models.TokenUser{}, errors.New("can't create a acces token")
	}

	refershToken, err := helper.GenerateRefreshToken(dataInsert)

	if err != nil {
		return &models.TokenUser{}, errors.New("can't create a Refersh token")

	}

	return &models.TokenUser{
		Users:        dataInsert,
		AccessToken:  accessToken,
		RefreshToken: refershToken,
	}, nil

}
