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
		message := fmt.Sprint("Register Account Failed : ", err.Error())
		Helper.NewCreateLogging(message, "log_RegisterUser_"+time.Now().Format("01-02-2006")+".log", "Error")

		response := Helper.APIResponse("Register Account Failed", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		message := fmt.Sprint("Register Account Failed : ", err.Error())
		Helper.NewCreateLogging(message, "log_RegisterUser_"+time.Now().Format("01-02-2006")+".log", "Error")

		response := Helper.APIResponse("Register Account Failed", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := Users.FormatUser(newUser, "TOKEN")
	response := Helper.APIResponse("Account has been registered", http.StatusOK, "Success", formatter)

	message := fmt.Sprint("Response Register User : ", response)
	Helper.NewCreateLogging(message, "log_RegisterUser_"+time.Now().Format("01-02-2006")+".log", "Info")

	c.JSON(http.StatusOK, response)
}
