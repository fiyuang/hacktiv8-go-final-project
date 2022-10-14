package middleware

import (
	_ "fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOnevalidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if _, err := strconv.Atoi(id); err != nil {
			// fmt.Println("Parameter False")
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Next()
	}
}
