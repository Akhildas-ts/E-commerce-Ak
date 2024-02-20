package routes

import (
	"ak/handlers"
	"ak/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	r.POST("/admin", handlers.AdminLogin)

	//MIDDLE WARE

	r.Use(middleware.AuthorizationMiddleware())
	{
		r.GET("/dashboard", handlers.DashBord)

		//PRODUCT
		product := r.Group("/products")
		{
			product.POST("", handlers.AddProduct)
			product.PUT("", handlers.UpdateProduct)
			product.DELETE("", handlers.DeleteProduct)
			product.GET("/:page", handlers.SeeAllProductToAdmin)
			product.POST("/add-product-offer", handlers.AddProductOffer)

		}

		// Categorry

		category := r.Group("/category")
		{
			category.POST("", handlers.AddCategory)
			category.PUT("", handlers.UpdateCategory)
			category.GET("", handlers.DeleteCategory)
			category.DELETE("/:page", handlers.GetAllCategory)
			category.POST("/add-category-offer", handlers.AddCategoryOffer)
		}

		// order

		order := r.Group("/order")
		{
			order.POST("/approve-order/:order_id", handlers.ApproveOrder)
			order.GET("/:page", handlers.GetOrderDetailsFromAdmin)
			order.DELETE("/cancel-order/:order_id", handlers.CancelOrderFromAdminSide)

		}

		offer := r.Group("/offer")
		{
			coupons := offer.Group("/coupons")

			{
				coupons.POST("", handlers.AddCoupon)
				coupons.GET("", handlers.GetCoupon)
				coupons.POST("/expire/:id", handlers.ExpireCoupon)
			}
		}

	}
	return r

}
