package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/AnonO6/geo-data-app/models"
	"github.com/AnonO6/geo-data-app/utils"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type AuthService struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewAuthService(db *gorm.DB, redisClient *redis.Client) *AuthService {
	return &AuthService{db, redisClient}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (as *AuthService) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate email format
	if !utils.IsValidEmail(req.Email) {
		utils.JSONError(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Validate password length
	if len(req.Password) < 6 {
		utils.JSONError(w, "Password must be at least 6 characters long", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.JSONError(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Create user
	user := models.User{
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := as.db.Create(&user).Error; err != nil {
		utils.JSONError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	utils.JSONSuccess(w, "User registered successfully", http.StatusCreated)
}

func (as *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Get user from the database
	var user models.User
	result := as.db.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.JSONError(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		utils.JSONError(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		utils.JSONError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	tokenString, err := utils.GenerateJWT(user.Email)
	if err != nil {
		utils.JSONError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Respond with success
	response := struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
		Token string `json:"token"`
	}{
		ID:    user.ID,
		Email: user.Email,
		Token: tokenString,
	}

	utils.JSONSuccess(w, response, http.StatusOK)
}

func (as *AuthService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	tokenString, err := utils.GetTokenFromHeader(r)
	if err != nil {
		utils.JSONError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Verify JWT token
	claims, err := utils.VerifyJWT(tokenString)
	if err != nil {
		utils.JSONError(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Parse user ID from token
	userID, err := strconv.Atoi(claims.Username)
	if err != nil {
		utils.JSONError(w, "Invalid user ID in token", http.StatusBadRequest)
		return
	}

	// Decode request body
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Get user from the database
	var user models.User
	if err := as.db.Where("id = ?", userID).First(&user).Error; err != nil {
		utils.JSONError(w, "User not found", http.StatusNotFound)
		return
	}

	// Update user details
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			utils.JSONError(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		user.Password = hashedPassword
	}

	if err := as.db.Save(&user).Error; err != nil {
		utils.JSONError(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	utils.JSONSuccess(w, "User updated successfully", http.StatusOK)
}
