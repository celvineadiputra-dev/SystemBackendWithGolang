package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"startup_be/Handler"
	"startup_be/Test"
	"startup_be/Users"
	"startup_be/auth"
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
	apiV1.POST("/avatars", userHandler.UploadAvatar)
	//END UPLOAD AVATAR

	router.Run()
}
