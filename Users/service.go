package Users

import (
	"fmt"
	"startup_be/Helper"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	message := fmt.Sprint("Body request register user : ", input)
	Helper.NewCreateLogging(message, "log_RegisterUser_"+time.Now().Format("01-02-2006")+".log", "Info")
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.OccupationId = input.OccupationId
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		message := fmt.Sprint("Save User To Database Failed : ", err.Error())
		Helper.NewCreateLogging(message, "log_RegisterUser_"+time.Now().Format("01-02-2006")+".log", "Error")
		return user, err
	}

	user.Password = string(passwordHash)
	user.RoleId = 1

	newUser, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}

	messageNew := fmt.Sprint("Save User To Database Success : ", newUser)
	Helper.NewCreateLogging(messageNew, "log_RegisterUser_"+time.Now().Format("01-02-2006")+".log", "Info")

	return newUser, nil
}
