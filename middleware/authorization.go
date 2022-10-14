package middleware

import (
	"fmt"
	"hacktiv8-go-final-project/database"
	"hacktiv8-go-final-project/models"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}
		getUser := db.First(&User, "id = ?", userID).Error

		fmt.Println(getUser)

		fmt.Println(userID)
		if c.Request.Method != "POST" {
			User := models.User{}
			err := db.First(&User, "id = ?", uint(userID)).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": "data doesn't exist",
				})
				return
			}

			// Data User Authorization
			if uint(User.Id) != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to access this data",
				})
				return
			}
		}

		c.Next()
	}
}
