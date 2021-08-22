package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"startup_be/Handler"
	"startup_be/Helper"
	"startup_be/Test"
	"startup_be/Users"
	"startup_be/auth"
	"startup_be/campaign"
	"strings"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()
	apiV1 := router.Group("api/v1")

	//TEST
	testRepository := Test.NewRepository(db)
	testService := Test.NewService(testRepository)
	testHandler := Handler.NewTestHandler(testService)
	//TEST

	//TEST CONNECTION TO DB
	apiV1.GET("/testConnect", testHandler.GetUser)
	//END CONNECTION TO DB

	// USERS
	authService := auth.NewService()

	userRepository := Users.NewRepository(db)
	userService := Users.NewService(userRepository)
	userHandler := Handler.NewUserHandler(userService, authService)
	// END USERS

	// Campaign
	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaigns, _ := campaignService.FindCampaigns(1)
	fmt.Println(len(campaigns))
	// End Campaign

	//REGISTER USER API
	apiV1.POST("/users", userHandler.RegisterUser) //Register User
	//END REGISTER USER API

	//LOGIN USER API
	apiV1.POST("/sessions", userHandler.Login)
	//END LOGIN USER API

	//CHECK EMAIL AVAILABLE
	apiV1.POST("/checkEmail", userHandler.CheckEmailAvailability)
	//END CHECK EMAIL AVAILABLE

	//UPLOAD AVATAR
	apiV1.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	//END UPLOAD AVATAR

	router.Run()
}

func authMiddleware(authService auth.Service, userService Users.Service) gin.HandlerFunc {
	return func (c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer"){
			response := Helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil{
			response := Helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid{
			response := Helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil{
			response := Helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}