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

	err = database.DB.Raw("select count(*) from users where blocked = false").Scan(&userDetails.UnBlockedUser).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	fmt.Println("hiiii")

	return userDetails, nil
}

func ChechOrderID(orderid string) (bool, error) {

	var count int

	err := database.DB.Raw("select count(*)from orders where order_id=?", orderid).Scan(&count).Error

	if err != nil {
		return false, err
	}
	if count == 0 {
		return false,errors.New("there is no record ")
	}

	return count > 0, nil

}

func ApproveOrder(orderID string) error {

	err := database.DB.Exec("update orders set shipment_status = 'order placed',approval = true where order_id = ?", orderID).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProductDetailsFromOrders(orderID string) ([]models.OrderProducts, error) {
	var orderProductDetails []models.OrderProducts
	if err := database.DB.Raw("select product_id,quantity from order_items where order_id = ?", orderID).Scan(&orderProductDetails).Error; err != nil {
		return []models.OrderProducts{}, err
	}
	return orderProductDetails, nil
}
