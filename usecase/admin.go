package usecase

import (
	"ak/domain"
	"ak/helper"
	"ak/models"
	"ak/repository"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(adminmodel models.AdminLogin) (domain.TokenAdmin, error) {

	// GETTING DETAILS FROM ADMIN WITH EMAIL

	AdminDetail, err := repository.AdminLogin(adminmodel)
	if err != nil {
		
		return domain.TokenAdmin{}, errors.New("given mail formate have")

	}

	// password := AdminDetail.Password

	if AdminDetail.Password == "" {
		return domain.TokenAdmin{}, errors.New("error from admin password")

	}

	err = bcrypt.CompareHashAndPassword([]byte(AdminDetail.Password), []byte(adminmodel.Password))

	if err != nil {

		return domain.TokenAdmin{},models.PasswordIsNotCorrect

	}

	var AdminDetailsResponse models.AdminDetailsResponse

	fmt.Println("ad", AdminDetailsResponse)

	err = copier.Copy(&AdminDetailsResponse,&AdminDetail) 

	if err!= nil{
		return domain.TokenAdmin{},err
	}


	tokenString, err := helper.GenerateTokenAdmin(AdminDetailsResponse)

	if err != nil {
		fmt.Println("error:3")
		return domain.TokenAdmin{}, errors.New("demo : repository adminlogin")

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

func ApproveOrder(orderid string) error {
	
	shipmentStatus, err := repository.GetShipmentStatus(orderid)
	if err != nil {
		return err
	}

	if shipmentStatus == "order placed" {
		return models.OrderIsAlreadyPlaced

	}
	ok, err := repository.ChechOrderID(orderid)

	if !ok {
		return err
	}

	
	

	if shipmentStatus == "pending" {
		return errors.New("the order is pending ,cannot approve it ")
	}

	if shipmentStatus == "cancelled" {
		return  models.OrderIsAlreadyCancelled
	}

	if shipmentStatus == "processing" {

		fmt.Println("reached")

		err := repository.ApproveOrder(orderid)

		if err != nil {
			return err
		}

	}

	return nil

}




func CancelOrderFromAdminSide(orderID string) error {

	_,err := repository.GetOrderDetailsByOrderID(orderID)

	if err != nil {
		return err
	}

	

	shipmentStatus,err := repository.GetShipmentStatus(orderID)

	if err!= nil {
		return err
	}

	fmt.Println("shipement status is ",shipmentStatus)

	if shipmentStatus == "cancelled"{
		return models.OrderIsAlreadyCancelled
	}

	orderProducts, err := repository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return err
	}

	err = repository.CancelOrders(orderID)
	if err != nil {
		return err
	}

	// update the quantity to products since the order is cancelled
	err = repository.UpdateQuantityOfProduct(orderProducts)
	if err != nil {
		return err
	}

	return nil

}
