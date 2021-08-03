package Handler

import (
	"net/http"
	"startup_be/Helper"
	"startup_be/Users"

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
		c.JSON(http.StatusBadRequest, nil)
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	formatter := Users.FormatUser(newUser, "TOKEN")
	response := Helper.APIResponse("Account has been registered", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}
