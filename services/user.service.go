package services

import (
	db "authService/db/repositories"
	model "authService/models"
	// "fmt"
)

type UserService interface{
	CreateUser() error
	GetById(id string) (error, *model.User)
}

type UserServiceImpl struct{
	UserRepository db.UserRepository
}

func (u *UserServiceImpl) CreateUser() error{
	err := u.UserRepository.Create()
	return err
}

func NewUserService(_userRepository db.UserRepository) UserService{
	return &UserServiceImpl{
		UserRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetById(id string) (error, *model.User){
	err, user := u.UserRepository.GetById(id)
	if err != nil {
		// fmt.Println("Error while fetching data from userRepo")
		return err, nil
	}
	
	return nil, user
}

