package usecase

import (
	"ak/config"
	"ak/helper"
	"ak/models"
	"ak/repository"
	"errors"
	"fmt"
)

func SendOTP(phone string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("cfg config")
		return errors.New("can not generate")
	}

	user, err := repository.FindUserByMobileNumber(phone)
	if err != nil {
		return errors.New("error with the server (find user by mobile number)")
	}

	if user == nil {
		return errors.New("phone number not exist ")

	}

	helper.TwilioSetUp(cfg.ACCOUNTSID, cfg.AUTHTOKEN)
	_, err = helper.TwilioSendOtp(phone, cfg.SERVICESSID)

	if err != nil {
		return errors.New("errors occured while generating otp")
	}

	return nil
}

func VerifyOTP(code models.VerifyData) (models.TokenUser, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return models.TokenUser{}, err
	}

	helper.TwilioSetUp(cfg.ACCOUNTSID, cfg.AUTHTOKEN)
	err = helper.TwilioVerifyOTP(cfg.SERVICESSID, code.Code, code.User.PhoneNumber)
	if err != nil {
		return models.TokenUser{}, err
	}

	UserDetails, err := repository.UserDetailsUsingPhone(code.User.PhoneNumber)
	if err != nil {
		return models.TokenUser{}, err
	}
	accessToken, err := helper.GenerateAccessToken(UserDetails)
	if err != nil {
		return models.TokenUser{}, err
	}

	refershToken, err := helper.GenerateRefreshToken(UserDetails)
	if err != nil {
		return models.TokenUser{}, err
	}

	var user models.SignupDetailResponse
	// err = copier.Copy(&user, &UserDetails)
	// if err != nil {
	// 	return models.TokenUser{}, err
	// }

	return models.TokenUser{
		Users:        user,
		AccessToken:  accessToken,
		RefreshToken: refershToken,
	}, nil

}
