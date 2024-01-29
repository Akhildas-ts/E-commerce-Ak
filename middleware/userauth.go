package middleware

import (
	"ak/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//retrive the jwt token from the header *****

		authheader := c.GetHeader("Authorization")

		tokenString := helper.GetTokenFromHeader(authheader)

		// VALIDATE THE TOKEN AND EXTRACT THE  USER ID

		if tokenString == "" {
			var err error
			tokenString, err = c.Cookie("Authorization")
			if err != nil {
				fmt.Println("error from token string")

				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		userId, userEmail, err := helper.ExtractUserIDFromToken(tokenString)
		fmt.Println("userId",userId,"userEmail",userEmail)

		if err != nil {
			fmt.Println("error from token ")
			fmt.Println("error is ",err)

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
