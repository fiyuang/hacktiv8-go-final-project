package controllers

import (
	"hacktiv8-go-final-project/helpers"
	"hacktiv8-go-final-project/models"
	"hacktiv8-go-final-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

//Kontrak Controllernya
type UserController interface {
	UserRegister(c *gin.Context)
	UserLogin(c *gin.Context)
	GetUserById(c *gin.Context)
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
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}
	res, err := controller.UserService.GetUserById(idint)
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

	c.JSON(http.StatusOK, gin.H{"data": res})
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
