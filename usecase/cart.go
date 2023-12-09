package usecase

import (
	"ak/models"
	"ak/repository"
	"errors"
	"fmt"
)

func AddToCart(product_id int, user_id int) (models.CartResponse,error) {

	QuantityOfProductInCart,err := repository.QuantityOfProductInCart(user_id,product_id)

	if err != nil {
		return models.CartResponse{},errors.New("dont have quantity")
	}

	productPrice,err :=repository.GetPriceOfProductFromID(product_id)

	if err != nil {
		return models.CartResponse{},errors.New("price of the product")
	}


	if QuantityOfProductInCart == 0 {
		err := repository.AddItemsIntoCart(user_id,product_id,1,productPrice)

		if err != nil {
			return models.CartResponse{},err
		}

	}

	cartDetails,err := repository.DisplayCart(user_id)

	if err !=nil {
		
		return models.CartResponse{},err
	}

	cartTotal, err := repository.GetTotalPrice(user_id)
	if err != nil {

		return models.CartResponse{}, err
	}

	fmt.Println("user name:",cartTotal.UserName)
	fmt.Println("totalprice",cartTotal.TotalPrice)
	fmt.Println("cart",cartDetails)

	
	return models.CartResponse{
		UserName: cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart: cartDetails,
		},nil
		
}


// REMOVE FROM CART ***

func RemoveFromCart(product_id int,user_id int) (models.CartResponse,error) {

	ok,err := repository.ProductExist(user_id,product_id)

	if err != nil {
		return models.CartResponse{},nil

	}

	if !ok {
		return models.CartResponse{},errors.New("product does not exist in the cart ")

	}

	var cartDetails struct {
		Quantity int
		TotalPrice float64
	}

	cartDetails,err = repository.GetQuantityAndProductDetails(user_id,product_id,cartDetails)

	if err != nil {

		return models.CartResponse{},err
	}
	cartDetails.Quantity =cartDetails.Quantity -1

	if  cartDetails.Quantity==0 {
		if err := repository.RemoveProductFromCart(user_id,product_id,);err != nil {
			return  models.CartResponse{},err
		}
	}

	updatecart,err := repository.CartAfterRemovalOfProduct(user_id)

	if err != nil {
		return models.CartResponse{},err
	}

	cartTotal,err := repository.GetTotalPrice(user_id)

	if err != nil {
		return models.CartResponse{},err
	}

	return models.CartResponse{
		UserName: cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart: updatecart,
	},nil



}

func DisplayCart(userid int ) (models.CartResponse,error) {

	cart,err := repository.DisplayCart(userid)

	if err != nil {

		return models.CartResponse{},err
	}

	cartTotal,err := repository.GetTotalPrice(userid)

	if err != nil {
		return models.CartResponse{},err
	}

	return models.CartResponse{
		UserName: cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart: cart,
	},nil


}