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
	r.Get("/get_by_id", ur.UserController.GetById)
	r.Get("/get_all", ur.UserController.GetAll)
}

func NewUserRouter(_userController controllers.UserController) Router{
	return &UserRouter{
		UserController: _userController,
	}
}