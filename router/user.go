package routes

import (
	"ak/handlers"
	"ak/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.UserLoginWithPassword)
	r.POST("/send-otp", handlers.SendOTP)
	r.POST("/verify-otp", handlers.VerifyOTP)

	
	
	r.GET("/payment", handlers.MakePaymentRazorPay)
	r.GET("/payment-success", handlers.VerifyPayment)

	r.Group("/users")
	r.Use(middleware.AuthMiddleware())
	{
		r.GET("/showaddres", handlers.GetAllAddress)
		r.POST("/addaddress", handlers.AddAddress)
		r.GET("/userdetails", handlers.UserDetails)

         
		// r.POST("/addbillingaddress",handlers.AddBillingAddress)

		// CART
		cart := r.Group("/cart")

		{
			cart.POST("/addtocart/:product_id", handlers.AddToCart)
			cart.DELETE("/removefromcart/:product_id", handlers.RemoveFromCart)
			cart.GET("/displaycart", handlers.DisplayCart)

		}
		product := r.Group("/product")
		{

			product.GET("/:page", handlers.SeeAllProductToUser)

		}

		r.GET("/checkout", handlers.CheckOut)
		r.GET("/place-order/:address_id/:payment", handlers.PlaceOrder)


		

		order := r.Group("/order") 
		{

			order.POST("", handlers.OrderItemsFromCart)
			order.GET("/:page", handlers.GetOrderDetails)
			order.POST("/:id", handlers.CancelOrder)
		}
		r.GET("/delivered/:order_id", handlers.OrderDelivered)
		r.GET("/cancel/:order_id", handlers.ReturnOrder)



		

		r.POST("/coupon/apply", handlers.ApplyCoupon)
		r.GET("/referral/apply",handlers.ApplyReferral)
		r.POST("/filter", handlers.FilterCategory)
		r.GET("/:page", handlers.SeeAllProductToUser)

		
		r.POST("/addimage",handlers.UploadImage)
		

	}

	return r

}
