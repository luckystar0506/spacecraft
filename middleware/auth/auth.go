package auth

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type ClaimKey string

func AuthorizationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/signup" || r.URL.Path == "/login" {
			next(w, r)
		} else {
			authHeader := strings.Split(r.Header.Get("Authorization"), " ")
			// Check if there are 2 parts and the first part is Bearer
			if len(authHeader) != 2 || authHeader[0] != "Bearer" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Get the Secret Key from environment variable
			secretKey := os.Getenv("SECRET_KEY")

			tokenString := authHeader[1]
			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})

			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			username := claims[("username")]

			if username == nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), ClaimKey("username"), username)
			next(w, r.WithContext(ctx))
		}
	}
}
