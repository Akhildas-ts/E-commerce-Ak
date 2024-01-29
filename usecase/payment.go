package usecase

import (
	"ak/config"
	"ak/models"
	"ak/repository"
	"errors"
	"fmt"

	"github.com/razorpay/razorpay-go"
)

func MakePaymentRazorPay(orderID string, userID int) (models.CombinedOrderDetails, string, error) {

	cfg, _ := config.LoadConfig()
	combinedOrderDetails, err := repository.GetOrderDetailsByOrderID(orderID)

	if err != nil {

		return models.CombinedOrderDetails{}, "", err
	}

	status, err := repository.CheckPaymentStatus(orderID)
	if err != nil {
		return models.CombinedOrderDetails{}, "", err
	}

	if status == "paid" {
		return models.CombinedOrderDetails{}, "", errors.New("alredy paid")

	}

	// err = repository.CheckOrderIdFromRazorPay(orderID)

	// if err != nil {
	// 	return models.CombinedOrderDetails{}, "", err
	// }

	client := razorpay.NewClient(cfg.KEY_ID_FOR_RAYZORPAY, cfg.SECRET_KEY_ID_FOR_RAYZORPAY)

	data := map[string]interface{}{
		"amount":   int(combinedOrderDetails.FinalPrice) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}

	body, err := client.Order.Create(data, nil)

	if err != nil {

		fmt.Println(err)
		return models.CombinedOrderDetails{}, "", err
	}

	razorPayOrderID := body["id"].(string)

	err = repository.AddRazorPayDetails(orderID, razorPayOrderID)
	if err != nil {
		return models.CombinedOrderDetails{}, "", err
	}

	return combinedOrderDetails, razorPayOrderID, nil

}
func SavePaymentDetails(paymentID string, orderID string) error {

	// to check whether the order is already paid
	status, err := repository.CheckPaymentStatus(orderID)
	if err != nil {
		return err
	}

	fmt.Println("status is ", status)

	if status == "not paid" {

		err = repository.UpdatePaymentDetails(orderID, paymentID)
		if err != nil {
			return err
		}

		err := repository.UpdateShipmentAndPaymentByOrderID("processing", "paid", orderID)
		if err != nil {
			return err
		}
		return nil
	}

	return models.AlreadyPaid

}
