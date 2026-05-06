package controllers

import (
	model "authService/models"
	"authService/services"
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	var payload struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := u.UserServices.CreateUser(payload.Username, payload.Email, payload.Password)
	if err != nil {
		http.Error(w, "User not created", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("user created successfully"))
}

func (u *UserController) LoginController(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, user, err := u.UserServices.LoginUser(payload.Email, payload.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Login successful",
		"token":   token,
		"user":    user,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (u *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	user, err := u.UserServices.GetById(id)
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
