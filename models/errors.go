package models

import "errors"

var (
	ErrEmailNotFound            = errors.New("email not found")
	CartEmpty                   = errors.New("cart empty")
	AddresNotFound              = errors.New("address not found")
	UserNotMatch                = errors.New("user not match")
	QuantityIsLessThanZero      = errors.New("quantity is less than zero")
	PriceIsLessThanZero         = errors.New("price is less than zero")
	ProductNameIsAlredyExist    = errors.New("product name is alredy exist")
	OrderIsAlreadyPlaced        = errors.New("order is already placed")
	OrderIsAlreadyCancelled     = errors.New("order is already cancelled")
	AlreadyPaid                 = errors.New("already paid")
	CannotReturn                = errors.New("cannot return after 7 days")
	AlreadyReturn               = errors.New("already return")
	AlreadyCancelled            = errors.New("already canncelled")
	AddresAlreadyExist          = errors.New("addres is already exist")
	PasswordIsNil               = errors.New("passwod is nil")
	PasswordIsNotCorrect        = errors.New("password is not correct")
	ShipmentStatusIsNotDeliverd = errors.New("shipment Status is not delivered")
	ThereIsNOCategory           = errors.New("there is no category ")
	OfferNameCantBeNil          = errors.New("offer name can't be nil")
	DiscountPriceIsLessThanZero = errors.New("Discount price is less than zero")
	OfferLimitGreater   = errors.New("offer limit must be greater than zero")
)
