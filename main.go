package main

import (
	"log"
	"net/http"
	"startup_be/Handler"
	"startup_be/Helper"
	"startup_be/Users"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()
	apiV1 := router.Group("api/v1")

	//TEST CONNECTION TO DB
	apiV1.GET("/testConect", testConnectToDb)
	//END CONNECTION TO DB

	//REGISTER USER API
	userRepository := Users.NewRepository(db)
	userService := Users.NewService(userRepository)
	userHandler := Handler.NewUserHandler(userService)

	apiV1.POST("/users", userHandler.RegisterUser) //Register User
	//END REGISTER USER API

	router.Run()
}

func testConnectToDb(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []Users.User
	db.Find(&users)

	Helper.NewCreateLogging("Test Connection to DB is Success", "log_testConnection_"+time.Now().Format("01-02-2006")+".log", "Info")

	c.JSON(http.StatusOK, users)
}
