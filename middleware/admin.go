package middleware

import (
	"ak/helper"
	"ak/response"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenheader := c.GetHeader("Authorization")
		fmt.Println(tokenheader, "this the token")
		if tokenheader == "" {
			response := response.ClientResponse(http.StatusUnauthorized, "no auth header provieded", nil, nil)
			c.JSON(http.StatusUnauthorized, response)

			fmt.Println(tokenheader, "this the token")

			c.Abort()
			return
		}

		splitted := strings.Split(tokenheader, " ")

		if len(splitted) != 2 {
			response := response.ClientResponse(http.StatusUnauthorized, "invalid token format", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return

		}

		tokenpart := splitted[1]
		tokenClaims, err := helper.ValidateToken(tokenpart)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		c.Set("tokenClaims", tokenClaims)
		c.Next()

	}
}
