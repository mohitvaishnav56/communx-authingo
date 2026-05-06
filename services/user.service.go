package services

import (
	db "authService/db/repositories"
	model "authService/models"
	"authService/utils"
	"fmt"
	// "fmt"
)

type UserService interface{
	CreateUser(username string, email string, password string) error
	LoginUser(email string, password string) (string, *model.User, error)
	GetById(id string) ( *model.User, error)
	GetAll() ([]*model.User, error)
}

type UserServiceImpl struct{
	UserRepository db.UserRepository
}

func (u *UserServiceImpl) CreateUser(username string, email string, password string) error{
	hashedPassword, error := utils.HashPassword(password)
	if error != nil {
		fmt.Println("Error while hashing the password")
	}
	err := u.UserRepository.Create(username, email, hashedPassword)
	return err
}

func NewUserService(_userRepository db.UserRepository) UserService{
	return &UserServiceImpl{
		UserRepository: _userRepository,
	}
}

func (u *UserServiceImpl) LoginUser(email string, password string) (string, *model.User, error) {
	user, err := u.UserRepository.GetByEmail(email)
	if err != nil {
		return "", nil, fmt.Errorf("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", nil, fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.Id)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate token")
	}

	return token, user, nil
}

func (u *UserServiceImpl) GetById(id string) ( *model.User, error){
	user, err:= u.UserRepository.GetById(id)
	if err != nil {
		// fmt.Println("Error while fetching data from userRepo")
		return nil, err
	}
	
	return user, nil
}

func(u *UserServiceImpl) GetAll() ([]*model.User, error){
	users, err := u.UserRepository.GetAll()
	if(err != nil){
		return nil, err
	}
	return users, nil
}