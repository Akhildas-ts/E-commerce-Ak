package handlers

import (
	"ak/models"
	"ak/response"
	"ak/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//cart

// @Summary Add to Cart
// @Description Add product to the cart using product id
// @Tags User Cart
// @Accept json
// @Produce json
// @Param id path string true "product-id"
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart/addtocart/{id} [post]
func AddToCart(c *gin.Context) {

	id := c.Param(models.Product_ID)

	productid, err := strconv.Atoi(id)

	if err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "product id is given wrong formate", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	user_id, _ := c.Get(models.User_id)

	cartResponse, err := usecase.AddToCart(productid, user_id.(int))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add product in cart ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	succRes := response.ClientResponse(http.StatusOK, "add proudct in carts ", cartResponse, nil)
	c.JSON(http.StatusOK, succRes)
}
// @Summary Remove product from cart
// @Description Remove specified product of quantity 1 from cart using product id
// @Tags User Cart
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart/removefromcart/{id} [delete]
func RemoveFromCart(c *gin.Context) {

	id := c.Param(models.Product_ID)

	product_id, err := strconv.Atoi(id)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "product id is not correct ", nil, err)
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	user_id, _ := c.Get(models.User_id)

	updateCart, err := usecase.RemoveFromCart(product_id, user_id.(int))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot remove prouduct from cart ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	succesRes := response.ClientResponse(200, "product removed successfully", updateCart, nil)
	c.JSON(200, succesRes)

}
// @Summary Display Cart
// @Description Display all products of the cart along with price of the product and grand total
// @Tags User Cart
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart/displaycart [get]
func DisplayCart(c *gin.Context) {

	userID, _ := c.Get(models.User_id)

	cart, err := usecase.DisplayCart(userID.(int))

	if err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "cannot display cart", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "carts items displayed successfully", cart, nil)
	c.JSON(http.StatusOK, succesRes)

}
// @Summary Apply coupon on Checkout Section
// @Description Add coupon to get discount on Checkout section
// @Tags User Checkout
// @Accept json
// @Produce json
// @Security Bearer
// @Param couponDetails body models.CouponAddUser true "Add coupon to order"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /coupon/apply [post]
func ApplyCoupon(c *gin.Context) {

	userID, _ := c.Get(models.User_id)

	var couponDetails models.CouponAddUser

	if err := c.ShouldBindJSON(&couponDetails); err != nil {

		errores := response.ClientResponse(http.StatusBadRequest, "could not blind the caption", nil, err.Error())

		c.JSON(http.StatusBadGateway, errores)
		return
	}

	err := usecase.ApplyCoupon(couponDetails.CouponName, userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "coupon could not  added,check it  ", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.SuccessClientResponse(http.StatusCreated, "Coupon added successfully")
	c.JSON(http.StatusCreated, successRes)
}
