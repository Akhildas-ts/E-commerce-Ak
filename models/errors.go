package models

import "errors"

var (
	ErrEmailNotFound = errors.New("email not found")
	CartEmpty        = errors.New("cart empty")
	AddresNotFound = errors.New("address not found")
	UserNotMatch = errors.New("user not match")
    QuantityIsLessThanZero = errors.New("quantity is less than zero")
	PriceIsLessThanZero = errors.New("price is less than zero")
	ProductNameIsAlredyExist = errors.New("product name is alredy exist")
	OrderIsAlreadyPlaced = errors.New("order is already placed")
	OrderIsAlreadyCancelled= errors.New("order is already cancelled")
	AlreadyPaid = errors.New("already paid")
	CannotReturn = errors.New("cannot return after 7 days")
	AlreadyReturn = errors.New("already return")
	AlreadyCancelled = errors.New("already canncelled")
	

)
