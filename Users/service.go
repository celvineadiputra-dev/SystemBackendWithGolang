package Users

import (
	"errors"
	"fmt"
	"startup_be/Helper"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
	IsEmailAvailable(input EmailUserInput) (bool, error)
	SaveAvatar(ID int, FileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
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

func (s *service) Login(input LoginUserInput) (User, error)  {
	message := fmt.Sprint("Body request login user : ", input)
	Helper.NewCreateLogging(message, "log_LoginUser_"+time.Now().Format("01-02-2006")+".log", "Info")
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil{
		message := fmt.Sprint("Login User Failed : ", err.Error())
		Helper.NewCreateLogging(message, "log_LoginUser_"+time.Now().Format("01-02-2006")+".log", "Error")
		return user, err
	}

	if user.ID == 0{
		message := "User not found with email : "+email
		return user, errors.New(message)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil{
		message := fmt.Sprint("Login User Failed : ", err.Error())
		Helper.NewCreateLogging(message, "log_LoginUser_"+time.Now().Format("01-02-2006")+".log", "Error")
		return user, err
	}

	messageNew := fmt.Sprint("Login User Success : ", user)
	Helper.NewCreateLogging(messageNew, "log_LoginUser_"+time.Now().Format("01-02-2006")+".log", "Info")

	return user, nil
}

func (s *service) IsEmailAvailable (input EmailUserInput) (bool, error){
	email := input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil{
		return false, err
	}

	if user.ID == 0 {
		// if email is not registered
		return true, nil
	}

	//if email is already registered
	return false, nil
}

func (s *service) SaveAvatar(ID int, FileLocation string) (User, error){
	user, err := s.repository.FindById(ID)
	if err != nil{
		message := fmt.Sprint("Save User Avatar Failed : ", err.Error())
		Helper.NewCreateLogging(message, "log_SaveAvatar_"+time.Now().Format("01-02-2006")+".log", "Error")
		return user, err
	}

	user.AvatarFileName = FileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil{
		message := fmt.Sprint("Save User Avatar Failed : ", err.Error())
		Helper.NewCreateLogging(message, "log_SaveAvatar_"+time.Now().Format("01-02-2006")+".log", "Error")
		return user, err
	}

	messageNew := fmt.Sprint("Save User Avatar Success : ", user)
	Helper.NewCreateLogging(messageNew, "log_SaveAvatar_"+time.Now().Format("01-02-2006")+".log", "Info")

	return updatedUser, nil
}

func (s *service) GetUserByID(ID int) (User, error){
	user, err := s.repository.FindById(ID)
	if err != nil{
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}

	return user, nil
}