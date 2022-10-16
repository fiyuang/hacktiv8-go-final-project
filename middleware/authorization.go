package middleware

import (
	"hacktiv8-go-final-project/database"
	"hacktiv8-go-final-project/models"
	"strconv"

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
		if c.Request.Method != "POST" {
			err := db.First(&User, "id = ?", userID).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": "Data doesn't exist",
				})
				return
			}

			// Data User Authorization
			if uint(User.Id) != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data",
				})
				return
			}
		}

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		// Data User Define
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}
		err := db.First(&User, "id = ?", userID).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "User Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if c.Request.Method != "POST" {
			photoId, err := strconv.Atoi(c.Param("photoId"))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request",
					"message": "invalid parameter",
				})
				return
			}

			// Data Photos Define
			Photos := models.Photo{}
			err = db.Select("user_id").First(&Photos, uint(photoId)).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": "Data doesn't exist",
				})
				return
			}

			if Photos.UserId != User.Id {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data",
				})
				return
			}

		}
		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		// Data User Define
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		User := models.User{}
		err := db.First(&User, "id = ?", userID).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "User Not Found",
				"message": "Data doesn't exist",
			})
			return
		}

		if c.Request.Method != "POST" {
			commentId, err := strconv.Atoi(c.Param("commentId"))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request",
					"message": "Invalid parameter",
				})
				return
			}

			// Data Comment Define
			Comment := models.Comment{}
			err = db.Select("user_id").First(&Comment, uint(commentId)).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": "Data doesn't exist",
				})
				return
			}

			if Comment.UserId != User.Id {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data",
				})
				return
			}

		}
		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		// Data User Define
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		User := models.User{}
		err := db.First(&User, "id = ?", userID).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "User Not Found",
				"message": "Data doesn't exist",
			})
			return
		}

		if c.Request.Method != "POST" {
			socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request",
					"message": "Invalid parameter",
				})
				return
			}

			// Data SocialMedia Define
			SocialMedia := models.SocialMedia{}
			err = db.Select("user_id").First(&SocialMedia, uint(socialMediaId)).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": "Data doesn't exist",
				})
				return
			}

			if SocialMedia.UserId != User.Id {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data",
				})
				return
			}

		}
		c.Next()
	}

}
