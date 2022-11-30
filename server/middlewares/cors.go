package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func CheckJSONApplicationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
        if c.Request.Method == "POST" {
		    header := c.GetHeader("content-type")
            if headerUpper := c.GetHeader("Content-Type"); headerUpper != "" {
                header = headerUpper
            }
		    if header != "application/json"{
			    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				    "message": "Content-Type now allowed",
			    })
		    }
        }
		c.Next()
	}
}
