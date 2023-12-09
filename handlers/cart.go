package handlers

import (
	"ak/response"
	"ak/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//cart

func AddToCart(c *gin.Context) {

	id := c.Param("product_id")

	productid, err := strconv.Atoi(id)

	if err != nil {

		errRes := response.ClientResponse(http.StatusBadGateway, "product id is given wrong formate", nil, err)
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	user_id, _ := c.Get("user_id")

	cartResponse, err := usecase.AddToCart(productid, user_id.(int))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "could not add product in cart ", nil, err)
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	succRes := response.ClientResponse(http.StatusOK, "add proudct in carts ", cartResponse, nil)
	c.JSON(http.StatusOK, succRes)
}

func RemoveFromCart(c *gin.Context) {

	id := c.Param("product_id")

	product_id, err := strconv.Atoi(id)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "product id is not correct ", nil, err)
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	user_id, _ := c.Get("user_id")

	updateCart, err := usecase.RemoveFromCart(product_id, user_id.(int))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "cannot remove prouduct from cart ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	succesRes := response.ClientResponse(200, "product removed successfully", updateCart, nil)
	c.JSON(200, succesRes)

}

func DisplayCart(c *gin.Context) {

	userID, _ := c.Get("user_id")

	cart, err := usecase.DisplayCart(userID.(int))

	if err != nil {

		errRes := response.ClientResponse(http.StatusBadGateway, "cannot display cart", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "carts items displayed successfully", cart, nil)
	c.JSON(http.StatusOK, succesRes)

}


