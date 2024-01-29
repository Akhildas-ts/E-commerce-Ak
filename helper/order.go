package helper

import (
	"ak/domain"
	"ak/models"
	"strconv"

	"github.com/google/uuid"
)

func CopyOrderDetails(orderDetails domain.Order, orderBody models.OrderIncoming) domain.Order {

	id := uuid.New().ID()
	str := strconv.Itoa(int(id))
	orderDetails.OrderId = str[:8]
	orderDetails.AddressID = orderBody.AddressID
	orderDetails.PaymentMethodID = orderBody.PaymentID
	orderDetails.UserID = int(orderBody.UserID)
	orderDetails.Approval = false
	orderDetails.ShipmentStatus = "processing"
	orderDetails.PaymentStatus = "not paid"

	return orderDetails

}
