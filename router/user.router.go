package router

import (
	"authService/controllers"
	"github.com/go-chi/chi/v5"
)

type UserRouter struct{
	UserController controllers.UserController
}

func (ur UserRouter) Register(r chi.Router){
	r.Post("/signup", ur.UserController.RegisterController)
}

func NewUserRouter(_userController controllers.UserController) Router{
	return &UserRouter{
		UserController: _userController,
	}
}