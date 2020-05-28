package middlewares

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yes5144/GoDevops/GoDevops_api/models"
	"github.com/yes5144/GoDevops/GoDevops_api/utils"
)

// JwtMiddleware xxx
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get tokenString from athorization header
		tokenString := c.GetHeader("Authorization")
		log.Println(tokenString)
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			utils.Fail(c, nil, "Authory not enough ")
			c.Abort()
			return
		}
		tokenString = tokenString[7:]

		// parseToen
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			utils.Fail(c, nil, "authory nt enough")
			c.Abort()
			return
		}
		//
		userId := claims.UserID
		DB := models.GetDB()
		var user models.User
		DB.First(&user, userId)

		// user exist
		if userId == 0 {
			utils.Fail(c, nil, "no authorization")
			c.Abort()
			return
		}

		// write into context
		c.Set("user", user)
		c.Next()
	}
}
