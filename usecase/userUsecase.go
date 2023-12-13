package usecase

import (
	"ak/helper"
	"ak/models"
	"ak/repository"
	"errors"
	"fmt"
	"net/mail"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
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

	fmt.Println("data inserted in signup :", dataInsert)

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

func UserLogged(user models.UserLogin) (*models.TokenUser, error) {

	_, err := mail.ParseAddress(user.Email)
	if err != nil {

		return &models.TokenUser{}, errors.New("EMAIL SHOULD BE CORRECT FORMAT ")

	}

	email, err := repository.CheckingEmailValidation(user.Email)

	if err != nil {
		return &models.TokenUser{}, errors.New("SERVER ERROR  from : checking-email-validation")

	}

	if email == nil {
		return &models.TokenUser{}, errors.New("EMAIL IS NOT AVAILABLE ")

	}

	userDetails, err := repository.FindUserDetailByEmail(user)
	if err != nil {
		return &models.TokenUser{}, err

	}

	// CHECKING THE HASSED PASSWORD

	err = bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(user.Passoword))
	if err != nil {
		fmt.Println("HEyyyy")
		return &models.TokenUser{}, errors.New("hassed password not matching")
	}
	var user_details models.SignupDetailResponse

	err = copier.Copy(&user_details, &userDetails)
	if err != nil {

		return &models.TokenUser{}, errors.New("eorro in copier")
	}

	// TOKEN....

	accessToken, err := helper.GenerateAccessToken(user_details)
	if err != nil {
		return &models.TokenUser{}, errors.New("could not create accesstoken due to internal error")
	}
	refreshToken, err := helper.GenerateRefreshToken(user_details)
	if err != nil {
		return &models.TokenUser{}, errors.New("counldnt create refresh token due to error")
	}

	return &models.TokenUser{
		Users:        user_details,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func GetAllAddress(userId int) (models.AddressInfoResponse, error) {

	addressInfo, err := repository.GetAllAddress(userId)

	if err != nil {
		return models.AddressInfoResponse{}, err
	}

	// if (models.AddressInfoResponse{}) != addressInfo{
	// 	return models.AddressInfoResponse{},errors.New("there is no record is avalable on database")

	// }

	return addressInfo, nil
}

func Addaddress(UserId int, address models.AddressInfo) error {

	if err := repository.AddAddress(UserId, address); err != nil {
		return err
	}

	return nil
}

func UserDetails(userid int) (models.UsersProfileDetails, error) {

	//  return repository.UserDetails(userid)

	// func UserDetails(userID int) (models.UsersProfileDetails, error) {
	models, error := repository.UserDetails(userid)

	fmt.Println("user details form repo :", models)
	return models, error

}

func CheckOut(userid int) (models.CheckoutDetails, error) {

	//LIST ALL ADDRESS  ADDED BY THE USER...

	allUserAddress, err := repository.GetAllAddresses(userid)

	fmt.Println("alluserAddres : from usecase:", allUserAddress)

	if err != nil {
		return models.CheckoutDetails{}, err

	}

	//GET ALL ITEMS FROM USER CARTS

	cartitems, err := repository.GetAllItemsFromCart(userid)

	fmt.Println("cartitems from repo:", cartitems)

	if err != nil {
		return models.CheckoutDetails{}, err
	}

	grandTotal, err := repository.GetTotalPrice(userid)

	if err != nil {

		return models.CheckoutDetails{}, err
	}

	// GET AVAILABLE THE PAYMENT OPTION <<<<

	// paymentDetails, err := repository.GetAllPaymentOption()

	// if err != nil {
	// 	return models.CheckoutDetails{}, err
	// }

	

	fmt.Println("modoels.checkoutdetails struct is :", models.CheckoutDetails{})

	fmt.Println("final price :", grandTotal.FinalPrice)
	return models.CheckoutDetails{
		AddressInfoResponse: allUserAddress,
		// Payment_Method: paymentDetails,
		Grand_Total:    grandTotal.TotalPrice,
		Total_Price:    grandTotal.FinalPrice,
		Cart:           cartitems,
		// Payment_Method: paymentDetails,
	}, nil

}
