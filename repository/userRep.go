package repository

import (
	"ak/database"
	"ak/domain"
	"ak/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CheckingEmailValidation(email string) (*domain.User, error) {

	var user domain.User

	result := database.DB.Where(&domain.User{Email: email}).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil

		}
		return nil, result.Error
	}
	fmt.Println("user details was :", user)
	return &user, nil

}

func CheckingPhoneExists(phone string) (*domain.User, error) {

	var user domain.User
	result := database.DB.Where(&domain.User{Phone: phone}).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}
	return &user, nil

}

func SignupInsert(user models.SignupDetail) (models.SignupDetailResponse, error) {

	var signupRes models.SignupDetailResponse

	err := database.DB.Raw("INSERT INTO users(firstname, lastname, email, phone, password) VALUES (?,?,?,?,?) RETURNING id,firstname,lastname,email,phone", user.FirstName, user.LastName, user.Email, user.Phone, user.Password).Scan(&signupRes).Error

	if err != nil {
		fmt.Println(err, "asd")
		return models.SignupDetailResponse{}, err
	}

	fmt.Println("signup inserted data's are :", signupRes)
	return signupRes, nil
}

func FindUserDetailByEmail(user models.UserLogin) (models.UserLoginResponse, error) {

	var UserDetails models.UserLoginResponse

	err := database.DB.Raw(
		`select * from users where email = ? and blocked = false`, user.Email).Scan(&UserDetails).Error

	if err != nil {
		return models.UserLoginResponse{}, errors.New("got an error fron ! searching users by email")

	}

	return UserDetails, nil
}
func FindUserByMobileNumber(phone string) (*domain.User, error) {
	var User domain.User

	result := database.DB.Where(&domain.User{Phone: phone}).First(&User)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &User, nil
}

func GetAllAddress(userId int) (models.AddressInfoResponse, error) {
	var addressInfo models.AddressInfoResponse

	if err := database.DB.Raw("select * from  addresses where user_id =?", userId).Scan(&addressInfo).Error; err != nil {
		return models.AddressInfoResponse{}, err
	}
	fmt.Println("get all address from repo :", addressInfo)

	return addressInfo, nil
}

func AddAddress(userId int, address models.AddressInfo) error {

	if err := database.DB.Raw("insert into addresses(user_id,name,house_name,street,city,state,pin) values(?,?,?,?,?,?,?)", userId, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin).Scan(&address).Error; err != nil {

		return err

	}
	return nil
}



func UserDetails(userid int) (models.UsersProfileDetails, error) {

	var userdetails models.UsersProfileDetails

	if database.DB == nil {
		fmt.Println("error from database")
		return models.UsersProfileDetails{}, errors.New("database connection is nil ")
	}

	err := database.DB.Raw("select users.firstname,users.lastname,users.email,users.phone from users  where users.id = ?", userid).Row().Scan(&userdetails.Firstname, &userdetails.Lastname, &userdetails.Email, &userdetails.Phone)
	if err != nil {
		return models.UsersProfileDetails{}, err
	}

	fmt.Println("userdeails:", userdetails)

	return userdetails, nil

}

func GetAllAddresses(userid int) ([]models.AddressInfoResponse, error) {

	var addressResponse []models.AddressInfoResponse

	err := database.DB.Raw(`select * from addresses where user_id=$1`, userid).Scan(&addressResponse).Error

	if addressResponse == nil {

		fmt.Println("there is no addres from : repo")
		return []models.AddressInfoResponse{}, err
	}

	if err != nil {

		return []models.AddressInfoResponse{}, errors.New("cound not find addresses from user")
	}

	fmt.Println("addres respo is :", addressResponse)

	return addressResponse, nil
}


func GetAllPaymentOption() ([]models.PaymentDetails,error)  {

	var paymentmethod []models.PaymentDetails

	err := database.DB.Raw("select * from  payment_methods").Scan(&paymentmethod).Error

	if err != nil {
		return []models.PaymentDetails{},err
	}

	return paymentmethod,nil
}

func GetReferralAndTotalAmount(userId int ) (float64,float64,error){

	var cartDetails struct{

		ReferralAmount float64
		totalCartAmount float64
	}


	err := database.DB.Raw("SELECT (SELECT referral_amount from referrals where user_id = ?)AS referral_amount,COALESCE(SUM(total_price),0)AS total_cart_amount from carts WHERE user_id= ?",userId,userId).Scan(&cartDetails).Error

	if err != nil{

		return 0.0,0.0,err
	}

	return cartDetails.ReferralAmount,cartDetails.totalCartAmount,nil
}

func UpdateSomethingBasedOnUserID(tableName string, columnName string, updateValue float64, userID int) error {

	err := database.DB.Exec("update "+tableName+" set "+columnName+" = ? where user_id = ?", updateValue, userID).Error
	if err != nil {
		database.DB.Rollback()
		return err
	}
	return nil

}

func CheckAddress(userid int )error {


	var addres models.AddressInfoResponse


	err := database.DB.Raw("select * from addresses where user_id = ?",userid).Scan(&addres).Error

	if err != nil {
		return err
	}

	// if (models.AddressInfoResponse{} != addres) {

	// 	return  models.AddresAlreadyExist
	// }
	return err
}
