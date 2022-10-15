package controllers

import (
	"errors"
	"hacktiv8-go-final-project/database"
	"hacktiv8-go-final-project/helpers"
	"hacktiv8-go-final-project/models"
	"hacktiv8-go-final-project/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	appJSON = "application/json"
)

//Kontrak Controllernya
type UserController interface {
	UserRegister(c *gin.Context)
	UserLogin(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userControllerImpl struct {
	UserService service.UserService
}

//Inisiasi struct dengan kontrak interface
func NewUserController(newUserService service.UserService) UserController {
	return &userControllerImpl{
		UserService: newUserService,
	}
}

func (controller *userControllerImpl) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}
	res, err := controller.UserService.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data One": res})
}

func (controller *userControllerImpl) UserRegister(c *gin.Context) {
	// Validate input
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create user
	res, err := controller.UserService.UserRegister(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 201,
		"data":   res,
	})
}

func (controller *userControllerImpl) UserLogin(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	getUser, err := controller.UserService.UserLogin(&User)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(getUser.Password), []byte(User.Password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(uint(User.Id), User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (controller *userControllerImpl) UpdateUser(c *gin.Context) {
	db := database.GetDB()

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userEmail := userData["email"].(string)

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := models.User{}
	error := db.First(&User, "email = ?", userEmail).Error
	if error != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
			return
		}
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := controller.UserService.UpdateUser(&input, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	userUpdate := models.UserUpdate{}
	userUpdate.Id = res.Id
	userUpdate.Username = res.Username
	userUpdate.Email = res.Email
	userUpdate.Age = res.Age
	userUpdate.UpdatedAt = res.UpdatedAt

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   userUpdate,
	})
}

func (controller *userControllerImpl) DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}
	res, err := controller.UserService.DeleteUser(uint(userId))
	_ = res

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
