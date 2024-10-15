package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain-text password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("Error hashing password: %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares a plaintext password with a hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// JSONSuccess sends a success response in JSON format
func JSONSuccess(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// JSONError sends an error response in JSON format
func JSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// JWT secret key (read from environment variables)
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Claims represents the structure of JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// VerifyJWT verifies the JWT token and returns the claims if valid
func VerifyJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// GenerateJWT generates a JWT token for a given username
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		logrus.Errorf("Error generating JWT token: %v", err)
		return "", err
	}
	return tokenString, nil
}

// IsValidEmail validates email format using regex
func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

// GetTokenFromHeader extracts the JWT token from the request header
func GetTokenFromHeader(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", errors.New("missing token")
	}
	return token, nil
}
