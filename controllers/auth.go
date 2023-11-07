package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"spacecraft/ent/user"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResult struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

// Signup function creates a new user in DB
func (sc *SpaceshipController) Signup(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 14)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = sc.client.User.
		Create().
		SetUsername(creds.Username).
		SetPassword(string(hash)). // Store the hashed password
		Save(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Result{Success: true})
}

// Login function verifies user credentials and generates a JWT
func (sc *SpaceshipController) Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	user, _ := sc.client.User.
		Query().
		Where(user.UsernameEQ(creds.Username)).
		Only(ctx)

	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
	})

	tokenString, err := token.SignedString([]byte("<Your_Signing_Key>"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(AuthResult{Success: true, Token: tokenString})
}
