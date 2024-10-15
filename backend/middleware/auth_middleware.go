package middleware

import (
	"net/http"
	"strings"

	"github.com/AnonO6/geo-data-app/utils"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		token := strings.Split(authHeader, "Bearer ")[1]
		if _, err := utils.VerifyJWT(token); err != nil {
			logrus.Errorf("Invalid JWT token: %v", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
