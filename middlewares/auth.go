package middlewares

import (
	"chant/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		userClaims, err := helper.AnalyseToken(token)
		if err != nil {
			c.Abort() // stop
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "Token check error",
			})
			return
		}
		c.Set("user_claims", userClaims)
		c.Next()
	}
}
