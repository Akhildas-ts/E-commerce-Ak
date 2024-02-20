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
		return &models.TokenUser{}, models.ErrEmailNotFound

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

func GetUserDataWithEmail(login models.UserLogin) (models.UserLoginResponse, error) {

	UserDetials, err := repository.FindUserDetailByEmail(login)
	if err != nil {
		return models.UserLoginResponse{}, errors.New("email does not exist on database ")

	}

	return UserDetials, nil

}

func GetAllAddress(userId int) (models.AddressInfoResponse, error) {

	addressInfo, err := repository.GetAllAddress(userId)

	if err != nil {
		return models.AddressInfoResponse{}, err
	}

	return addressInfo, nil
}

func Addaddress(UserId int, address models.AddressInfo) error {

	err := repository.CheckAddress(UserId)

	if err != nil {
		return err
	}

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

	if err != nil {
	
		return models.CheckoutDetails{}, err

	}

	//GET ALL ITEMS FROM USER CARTS

	cartitems, err := repository.GetAllItemsFromCart(userid)

	if err != nil {
		return models.CheckoutDetails{}, err
	}

	grandTotal, err := repository.GetTotalPrice(userid)

	if err != nil {

		return models.CheckoutDetails{}, err
	}

	var discountApplied []string

	err = repository.DiscountReason(userid, "used_coupons", "COUPON APPLIED", &discountApplied)
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	// GET AVAILABLE THE PAYMENT OPTION <<<<

	paymentDetails, err := repository.GetAllPaymentOption()

	if err != nil {
		return models.CheckoutDetails{}, err
	}

	return models.CheckoutDetails{
		AddressInfoResponse: allUserAddress,
		Grand_Total:         grandTotal.TotalPrice,
		Total_Price:         grandTotal.FinalPrice,
		Cart:                cartitems,
		Payment_Method:      paymentDetails,
		DiscountReason:      discountApplied,
	}, nil

}

func ApplyReferral(userid int) (string, error) {

	exist, err := repository.CartExist(userid)

	if err != nil {

		return "", err

	}

	if !exist {

		return "", models.CartEmpty
	}

	referralAmount, totalCartAmount, err := repository.GetReferralAndTotalAmount(userid)
	fmt.Println("totatl cart amount ", totalCartAmount, "refferal amount", referralAmount)

	if err != nil {

		return "", err
	}

	if totalCartAmount > referralAmount {
		totalCartAmount = totalCartAmount - referralAmount

		referralAmount = 0

	} else {

		referralAmount = referralAmount - totalCartAmount

		totalCartAmount = 0
	}

	err = repository.UpdateSomethingBasedOnUserID("referrals", "referral_amount", referralAmount, userid)

	if err != nil {
		return "", err
	}

	err = repository.UpdateSomethingBasedOnUserID("carts", "total_price", totalCartAmount, userid)

	fmt.Println("totatl cart amount ", totalCartAmount, "refferal amount", referralAmount)

	if err != nil {

		return "", err
	}

	return "", nil

}
