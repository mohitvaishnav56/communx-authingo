package controllers

import (
	"authService/services"
	"net/http"
)

type UserController struct{
	UserServices services.UserService
}

func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		UserServices: _userService,
	}
}

func (u *UserController) RegisterController(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello We are here"))
}