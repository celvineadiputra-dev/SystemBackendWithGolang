package Handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"startup_be/Helper"
	"startup_be/Test"
	"time"
)

type testHandler struct {
	testService Test.Service
}

func NewTestHandler(testService Test.Service)  *testHandler{
	return &testHandler{testService}
}

func (h *testHandler) GetUser(c *gin.Context) {
	user, err := h.testService.GetUserTest()
	if err != nil{
		errorMessage := gin.H{"errors" : err.Error()}
		response := Helper.APIResponse("TEST CONNECT DATABASE FAILED", http.StatusUnprocessableEntity, "Error", errorMessage)

		message := fmt.Sprint("TEST CONNECT DATABASE FAILED : ", response)
		Helper.NewCreateLogging(message, "TEST_CONNECT_DB"+time.Now().Format("01-02-2006")+".log", "Error")

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	c.JSON(http.StatusOK, user)
}