package controllers

import (
	"authService/services"
	"encoding/json"
	// "fmt"
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
	err := u.UserServices.CreateUser()
	if err != nil {
		// fmt.Println("Error while fetching data from userService")
		http.Error(w, "User not created", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("user created successfully"))
}
func (u *UserController) GetById(w http.ResponseWriter, r *http.Request){
	err, user := u.UserServices.GetById("4a6a9fb5-4897-11f1-8e38-7c2a318083c4")
	if err != nil {
		// fmt.Println("Error while fetching data from userService")
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	
}