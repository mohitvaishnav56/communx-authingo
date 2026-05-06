package services

import (
	db "authService/db/repositories"
	model "authService/models"
	// "fmt"
)

type UserService interface{
	CreateUser() error
	GetById(id string) ( *model.User, error)
	GetAll() ([]*model.User, error)
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