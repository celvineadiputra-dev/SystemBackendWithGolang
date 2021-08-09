package Handler

import (
	"fmt"
	"net/http"
	"startup_be/Helper"
	"startup_be/Users"
	"time"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService Users.Service
}

func NewUserHandler(userService Users.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input Users.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := Helper.FormatValidationError(err)
		errorMessage := gin.H{"erros": errors}

		message := fmt.Sprint("Register Account Failed : ", errorMessage)
		Helper.NewCreateLogging(message, "log_RegisterUser_"+time.Now().Format("01-02-2006")+".log", "Error")

		response := Helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := Helper.APIResponse("Register Account Failed", http.StatusBadRequest, "Error", err.Error())

		message := fmt.Sprint("Register Account Failed : ", response)
		Helper.NewCreateLogging(message, "log_RegisterUser_"+time.Now().Format("01-02-2006")+".log", "Error")

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := Users.FormatUser(newUser, "TOKEN")
	response := Helper.APIResponse("Account has been registered", http.StatusOK, "Success", formatter)

	message := fmt.Sprint("Response Register User : ", response)
	Helper.NewCreateLogging(message, "log_RegisterUser_"+time.Now().Format("01-02-2006")+".log", "Info")

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context){
	var input Users.LoginUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil{
		errors := Helper.FormatValidationError(err)
		errorMessage := gin.H{"errors" : errors}

		message := fmt.Sprint("Login User Failed : ", errorMessage)
		Helper.NewCreateLogging(message, "log_LoginUser_"+time.Now().Format("01-02-2006")+".log", "Error")

		response := Helper.APIResponse("Login User Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedUser, err := h.userService.Login(input)
	if err != nil{
		errorMessage := gin.H{"errors" : err.Error()}
		response := Helper.APIResponse("Login User Failed", http.StatusUnprocessableEntity, "Error", errorMessage)

		message := fmt.Sprint("Login User Failed : ", response)
		Helper.NewCreateLogging(message, "log_LoginUser_"+time.Now().Format("01-02-2006")+".log", "Error")

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := Users.FormatUser(loggedUser, "TOKEN")
	response := Helper.APIResponse("Login Success", http.StatusOK, "Success", formatter)

	message := fmt.Sprint("Response Register User : ", response)
	Helper.NewCreateLogging(message, "log_LoginUser_"+time.Now().Format("01-02-2006")+".log", "Info")

	c.JSON(http.StatusOK, response)
}