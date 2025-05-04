package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"layersapi/entities/dto"
	"layersapi/services"
	"log"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (u UserController) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	resData, err := u.userService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	res, err := json.Marshal(resData)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (u UserController) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if len(id) == 0 {
		http.Error(w, "id cannot be empty",
			http.StatusBadRequest)
		return
	}

	user, err := u.userService.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(user)
}

func (u UserController) CreateUserHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed",
			http.StatusMethodNotAllowed)
		return
	}

	var user dto.CreateUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if len(user.Name) == 0 {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	err = u.userService.Create(user)
	if err != nil {
		if errors.Is(err, errors.New("invalid email address")) ||
			errors.Is(err, errors.New("name must only contain alphabetic characters")) {
			http.Error(w, err.Error(),
				http.StatusBadRequest)
			return
		}
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func (u UserController) UpdateUserHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	vars := mux.Vars(r)
	id := vars["id"]

	if len(id) == 0 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var user dto.UpdateUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	if len(user.Name) == 0 {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "email is required", http.StatusBadRequest)
	}

	err = u.userService.Update(id, user)
	if err != nil {
		if errors.Is(err, errors.New("name must be only contain alphabetic characters")) ||
			errors.Is(err, errors.New("invalid email address")) {
			http.Error(w, "failed to update", http.StatusBadRequest)
			return
		}
		log.Println("Error al actualizar usuario:", err)
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user updated successfully"))
}

func (u UserController) DeleteUserHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	if len(id) == 0 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	err := u.userService.Delete(id)
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user deleted successfully"))
}
