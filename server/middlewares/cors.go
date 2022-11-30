package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func CheckJSONApplicationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Content-Type")
		if header != "application/json"{
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Content-Type now allowed",
			})
		}
		c.Next()
	}
}
