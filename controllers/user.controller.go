package controllers

import (
	model "authService/models"
	"authService/services"
	"encoding/json"
	// "fmt"
	"net/http"
)

type UserController struct {
	UserServices services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserServices: _userService,
	}
}

func (u *UserController) RegisterController(w http.ResponseWriter, r *http.Request) {
	err := u.UserServices.CreateUser()
	if err != nil {
		// fmt.Println("Error while fetching data from userService")
		http.Error(w, "User not created", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("user created successfully"))
}

func (u *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	user, err := u.UserServices.GetById("4a6a9fb5-4897-11f1-8e38-7c2a318083c4")
	if err != nil {
		// fmt.Println("Error while fetching data from userService")
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (u *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserServices.GetAll()
	if err != nil {
		http.Error(w, "error while fetching all the users", http.StatusInternalServerError)
		return
	}
	type CustomResponse struct {
		Status  int
		Message string
		Data    []*model.User
	}

	response := CustomResponse{
		Status:  http.StatusOK,
		Message: "data found successfully",
		Data:    users,
	}
	// fmt.Println(response)
	
	w.Header().Set("Content-Type", "application/json")
	// fmt.Println( json.NewEncoder(w).Encode(response))
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"status":"error","message":"Failed to encode JSON"}`, http.StatusInternalServerError)
	}
	
}
