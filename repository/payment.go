package repository

import (
	"ak/database"
)

func CheckPaymentStatus(orderID string) (string, error) {

	var paymentStatus string
	err := database.DB.Raw("select payment_status from orders where order_id = ?", orderID).Scan(&paymentStatus).Error
	if err != nil {
		return "", err
	}
	return paymentStatus, nil
}

func UpdatePaymentDetails(orderID string, paymentID string) error {

	err := database.DB.Exec("update razer_pays set payment_id = ? where order_id = ?", paymentID, orderID).Error
	if err != nil {
		return err
	}
	return nil

}

func UpdateShipmentAndPaymentByOrderID(shipmentStatus string, paymentStatus string, orderID string) error {

	err := database.DB.Exec("update orders set payment_status = ?, shipment_status = ? where order_id = ?", paymentStatus, shipmentStatus, orderID).Error
	if err != nil {
		return err
	}

	return nil

}

// func CheckOrderIdFromRazorPay(order_id string) error {

// 	var demoOrderId string

// 	err := database.DB.Raw("select order_id from razer_pays where order_id = ?", order_id).Scan(&demoOrderId).Error

// 	if err != nil {

// 		return err
// 	}

// 	if demoOrderId != "" {

// 		return errors.New("the order is done alredy ")
// 	}

// 	return nil

// }
