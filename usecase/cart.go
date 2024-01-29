package usecase

import (
	"ak/models"
	"ak/repository"
	"errors"
	"fmt"
)

func AddToCart(product_id int, user_id int) (models.CartResponse, error) {

	QuantityOfProductInCart, err := repository.QuantityOfProductInCart(user_id, product_id)

	const minQuantityAllowed = 0

	if err != nil {
		return models.CartResponse{}, errors.New("dont have quantity")
	}

	if QuantityOfProductInCart < minQuantityAllowed {
		return models.CartResponse{}, errors.New("quantity in cart is less than allowed minimum")
	}

	productPrice, err := repository.GetPriceOfProductFromID(product_id)

	if err != nil {
		return models.CartResponse{}, errors.New("price of the product")
	}

	if QuantityOfProductInCart == 0 {
		err := repository.AddItemsIntoCart(user_id, product_id, 1, productPrice)

		if err != nil {
			return models.CartResponse{}, err
		}

	} else if QuantityOfProductInCart > minQuantityAllowed {

		err := repository.IncreaseQuantiyInCart(user_id, product_id, 1, productPrice)
		if err != nil {

			return models.CartResponse{}, err
		}

	}

	cartDetails, err := repository.DisplayCart(user_id)

	if err != nil {

		return models.CartResponse{}, err
	}

	cartTotal, err := repository.GetTotalPrice(user_id)
	if err != nil {

		return models.CartResponse{}, err
	}



	return models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		FinalPrice: cartTotal.FinalPrice,
		Cart:       cartDetails,
	}, nil

}

// REMOVE FROM CART ***

func RemoveFromCart(product_id int, user_id int) (models.CartResponse, error) {

	ok, err := repository.ProductExist(user_id, product_id)

	if err != nil {
		return models.CartResponse{}, nil

	}

	if !ok {
		return models.CartResponse{}, errors.New("product does not exist in the cart ")

	}

	var cartDetails struct {
		Quantity   int
		TotalPrice float64
	}

	cartDetails, err = repository.GetQuantityAndProductDetails(user_id, product_id, cartDetails)

	if err != nil {

		return models.CartResponse{}, err
	}
	cartDetails.Quantity = cartDetails.Quantity - 1

	if cartDetails.Quantity == 0 {
		if err := repository.RemoveProductFromCart(user_id, product_id); err != nil {
			return models.CartResponse{}, err
		}
	}

	if cartDetails.Quantity != 0 {

		productPrice,err := repository.CheckProductPrice(product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
		fmt.Println("product Price",productPrice)
		cartDetails.TotalPrice = cartDetails.TotalPrice-productPrice

		if err = repository.UpdateCartDetails(cartDetails, user_id, product_id); err != nil {
			return models.CartResponse{}, err
		}
	}

	updatecart, err := repository.CartAfterRemovalOfProduct(user_id)

	if err != nil {
		return models.CartResponse{}, err
	}

	cartTotal, err := repository.GetTotalPrice(user_id)

	if err != nil {
		return models.CartResponse{}, err
	}

	return models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		FinalPrice: cartTotal.FinalPrice,
		Cart:       updatecart,
	}, nil

}

func DisplayCart(userid int) (models.CartResponse, error) {

	cart, err := repository.DisplayCart(userid)

	if err != nil {

		return models.CartResponse{}, err
	}

	cartTotal, err := repository.GetTotalPrice(userid)

	if err != nil {
		return models.CartResponse{}, err
	}

	return models.CartResponse{
		UserName:   cartTotal.UserName,
		FinalPrice: cartTotal.FinalPrice,
		TotalPrice: cartTotal.TotalPrice,

		Cart: cart,
	}, nil

}

func ApplyCoupon(coupon string, userid int) error {
	fmt.Println("User id ", userid)

	cartExist, err := repository.DoesCartExist(userid)

	if err != nil {

		return err
	}

	if !cartExist {

		return errors.New("cart empty can't  apply coupon ")
	}

	couponExist, err := repository.CouponExist(coupon)

	if err != nil {

		return err
	}

	if !couponExist {

		return errors.New("coupon does not exist")
	}

	couponValidity, err := repository.CouponValidity(coupon)

	if err != nil {

		return err
	}

	if !couponValidity {
		return errors.New("coupon expired")
	}

	minDiscountPrice, err := repository.GetCouponMinimumAmount(coupon)

	if err != nil {

		return err
	}

	totalPriceFromCarts, err := repository.GetTotalPriceFromCart(userid)

	if err != nil {

		return err
	}

	if totalPriceFromCarts < minDiscountPrice {

		return errors.New("coupon cannot be added as the total amount is less than minimum amount for coupon")
	}

	userAlreadyUsed, err := repository.DidUserAlreadyUsedThisCoupon(coupon, userid)

	if err != nil {

		return err
	}

	if err != nil {

		return err
	}

	if userAlreadyUsed {

		return errors.New("user already used this coupon")
	}

	couponStatus, err := repository.UpdateUsedCoupon(coupon, userid)

	if err != nil {

		return nil
	}

	if couponStatus {
		totalPriceFromCarts = totalPriceFromCarts - minDiscountPrice

		fmt.Println(totalPriceFromCarts)

		return nil
	}

	return errors.New("could not add the couupon")

}
