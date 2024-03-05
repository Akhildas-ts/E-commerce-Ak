package middleware

import (
	"ak/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Request Headerss:", c.Request.Header)

		//retrive the jwt token from the header *****

		authheader := c.GetHeader("Authorization")
		

		tokenString := helper.GetTokenFromHeader(authheader)

		// VALIDATE THE TOKEN AND EXTRACT THE  USER ID

		if tokenString == "" {
			var err error
			tokenString, err = c.Cookie("Authorization")
			if err != nil {
				

				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		userId, userEmail, err := helper.ExtractUserIDFromToken(tokenString)
		fmt.Println("userId", userId, "userEmail", userEmail)

		if err != nil {
		

			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//ADD USER ID ON THE GIN CONTEXT ******

		c.Set("user_id", userId)
		c.Set("user_email", userEmail)

		// CALL THE NEXT HANDLER ****

		c.Next()

	}
}
