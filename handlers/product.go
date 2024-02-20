package handlers

import (
	"ak/config"
	"ak/domain"
	"ak/models"
	"ak/response"
	s3bucket "ak/s3-bucket"
	"ak/usecase"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//ADD PRODUCT

// @Summary Add Product
// @Description Add product from Admin
// @Tags Admin Products
// @Accept json
// @Produce json
// @Security Bearer
// @Param product body domain.Products true "Product object to be added"
// @Success 200 {object} response.Response{data=domain.Products} "Successful response"
// @Failure 400 {object} response.Response{} "Bad Request"
// @Failure 404 {object} response.Response{} "Not Found"
// @Failure 500 {object} response.Response{} "Internal Server Error"
// @Router /admin/products [post]
func AddProduct(c *gin.Context) {

	var product domain.Products

	if err := c.ShouldBindJSON(&product); err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "field provide are wrong formate (product )", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	err := validator.New().Struct(product)
	if err != nil {
		errRes := response.ClientResponse(http.StatusNotFound, "Product structer is not correct ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return

	}

	ProductResponse, err := usecase.AddProduct(product)

	if errors.Is(err, models.QuantityIsLessThanZero) || errors.Is(err, models.PriceIsLessThanZero) || errors.Is(err, models.ProductNameIsAlredyExist) {

		errRes := response.ClientResponse(http.StatusBadRequest, "request is not correct", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	if err != nil {
		errRes := response.ClientResponse(http.StatusNotFound, "can't add the product", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, " Product is addeed ", ProductResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

//UPDATE THE PRODUCT*****

// @Summary Update Products
// @Description Update product from Admin side
// @Tags Admin Products
// @Accept json
// @Produce json
// @Security Bearer
// @Param p  body models.ProductUpdate true "Product object to be Updated"
// @Success 200 {object} response.Response{data=domain.Products} "Successful response"
// @Failure 400 {object} response.Response{} "Bad Request"
// @Failure 404 {object} response.Response{} "Not Found"
// @Failure 500 {object} response.Response{} "Internal Server Error"
// @Router /admin/products [put]
func UpdateProduct(c *gin.Context) {

	var p models.ProductUpdate

	if err := c.ShouldBindJSON(&p); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "field formate provide wrong formate", nil, err)

		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	updating, err := usecase.UpdateProduct(p.ProductId, p.Quantity)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not update the product ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully updating the product ", updating, nil)
	c.JSON(http.StatusOK, successRes)

}

/// DELET THE PROUDCT ****

// @Summary Delete Product
// @Description Delete product from Admin side
// @Tags Admin Products
// @Accept json
// @Produce json
// @Security Bearer
// @Param id  query string true "Product ID for deletion"
// @Success 200 {object} response.Response{data=domain.Products} "Successful response"
// @Failure 400 {object} response.Response{} "Bad Request"
// @Failure 404 {object} response.Response{} "Not Found"
// @Failure 500 {object} response.Response{} "Internal Server Error"
// @Router /admin/products [delete]
func DeleteProduct(c *gin.Context) {
	productID := c.Query(models.ID)

	if productID == "" {

		errRes := response.ClientResponse(http.StatusNotFound, "there is no product id", nil, fmt.Errorf("no ID provided"))
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	err := usecase.DeleteProduct(productID)
	if err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "field provide are wrong formate ", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	succRes := response.SuccessClientResponse(http.StatusOK, "Product was deleted succesfully")
	fmt.Println(succRes)
	c.JSON(http.StatusOK, succRes)

}

//SEE ALL PRODUCT TO ADMIN

// @Summary GET products DETAILS FROM ADMIN
// @Description Products details
// @Tags Admin Order Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param page path int true "page number"
// @Param count query int true "count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products/{page} [get]
func SeeAllProductToAdmin(c *gin.Context) {

	pagestr := c.Param(models.Page)
	page, err := strconv.Atoi(pagestr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	count, err := strconv.Atoi(c.Query(models.Count))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := usecase.ShowAllProducts(page, count)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not retrieve products", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Retrieved all products", products, nil)
	c.JSON(http.StatusOK, successRes)

}

// FILTER CATEGORY
// @Summary Show Products of specified category
// @Description Show all the Products belonging to a specified category
// @Tags User Product
// @Accept json
// @Security Bearer
// @Produce json
// @Param data body map[string]int true "Category IDs and quantities"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /filter [post]
func FilterCategory(c *gin.Context) {
	var data map[string]int
	if err := c.ShouldBindJSON(&data); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	productCategory, err := usecase.FilterCategory(data)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve products by category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully filtered the category", productCategory, nil)
	c.JSON(http.StatusOK, successRes)
}

// SEE ALL PRODUCT TO USER

// @Summary Get Products Details to users
// @Description Retrieve all product Details with pagination to users
// @Tags User Product
// @Accept json
// @Produce json
// @Security Bearer
// @Param page path int true "Page number"
// @Param count query int true "Page Count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /product/{page} [get]
func SeeAllProductToUser(c *gin.Context) {

	pagestr := c.Param(models.Page)
	page, err := strconv.Atoi(pagestr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	count, err := strconv.Atoi(c.Query(models.Count))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := usecase.ShowAllProducts(page, count)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not retrieve products", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Retrieved all products", products, nil)
	c.JSON(http.StatusOK, successRes)

}

func UploadImage(c *gin.Context) {

	cfg, err := config.LoadConfig()
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cfg  detials have error", nil, err)
		c.JSON(http.StatusInternalServerError, errRes)
		return

	}

	var awsInfo s3bucket.AWSConfig
	awsInfo.AccessKeyID = cfg.AWS_ACCESS_KEY_ID
	awsInfo.AccessKeySecret = cfg.AWS_SECRET_ACCESS_KEY
	awsInfo.BaseURL = cfg.BASE_URL
	awsInfo.BucketName = "akhils3-bucket"
	awsInfo.Region = cfg.AWS_REGION
	awsInfo.UploadTimeout = 10

	session := s3bucket.CreateSession(awsInfo)
	s3bucket.CreateS3Session(session)

	filename := "product.jpeg"
	filepath := "./image/"

	Imageurl, err := s3bucket.UploadObject(awsInfo.BucketName, filepath, filename, session, awsInfo)

	if err != nil {

		errRes := response.ClientResponse(http.StatusInternalServerError, "error in uploaded image", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	success := response.ClientResponse(http.StatusOK, "image uploaded succesfully", Imageurl, nil)

	c.JSON(http.StatusInternalServerError, success)
}
