package controllers

import (
	"log"
	"net/http"

	"github.com/AnonO6/geo-data-app/services"
	"github.com/AnonO6/geo-data-app/utils"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService}
}

// Register handles the registration of a new user
func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("Registering user")
	if r.Method != http.MethodPost {
		utils.JSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ac.authService.Register(w, r) // Use AuthService handler
}

// Login handles user login and token generation
func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.JSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ac.authService.Login(w, r) // Use AuthService handler
}

// UpdateUser handles updating user details (requires JWT token)
func (ac *AuthController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.JSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ac.authService.UpdateUser(w, r) // Use AuthService handler for user updates
}
