package main

import (
	"log"
	"net/http"
	"startup_be/Users"

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
	router.GET("/testConnecttoDb", testConnectToDb)

	// testRepository := Test.NewRepositoryTest(db)
	// var users []Users.User
	// router.GET("/testConnecttoDb", testRepository.Get(users))
	// router.Run()

	userRepository := Users.NewRepository(db)
	userService := Users.NewService(userRepository)

	userInput := Users.RegisterUserInput{}
	userInput.Name = "TEST SIMPAN DATA"
	userInput.Email = "TEST@MAIL.co.id"
	userInput.OccupationId = 1
	userInput.Password = "password"

	userService.RegisterUser(userInput)

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

	c.JSON(http.StatusOK, users)
}
