package services

import db "authService/db/repositories"

type UserService interface{
	CreateUser() error
}

type UserServiceImpl struct{
	UserRepository db.UserRepository
}

func (u *UserServiceImpl) CreateUser() error{
	return nil
}

func NewUserService(_userRepository db.UserRepository) UserService{
	return &UserServiceImpl{
		UserRepository: _userRepository,
	}
}