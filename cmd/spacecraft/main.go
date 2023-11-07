package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"spacecraft/controllers"
	"spacecraft/ent"
	"spacecraft/middleware/auth"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser, exists := os.LookupEnv("DATABASE_USER")
	if !exists {
		log.Fatal("DATABASE_USER not found in .env file")
	}

	dbPassword, exists := os.LookupEnv("DATABASE_PASSWORD")
	if !exists {
		log.Fatal("DATABASE_PASSWORD not found in .env file")
	}

	dbName, exists := os.LookupEnv("DATABASE_NAME")
	if !exists {
		log.Fatal("DATABASE_NAME not found in .env file")
	}

	dbHost, exists := os.LookupEnv("DATABASE_HOST")
	if !exists {
		log.Fatal("DATABASE_HOST not found in .env file")
	}

	dbPort, exists := os.LookupEnv("DATABASE_PORT")
	if !exists {
		log.Fatal("DATABASE_PORT not found in .env file")
	}

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	client, err := ent.Open("mysql", dbURL)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	controller := controllers.NewSpaceshipController(client)

	r := mux.NewRouter()

	r.HandleFunc("/spaceships", controller.GetSpaceships).Methods("GET")
	r.HandleFunc("/spaceships/{id}", controller.GetSpaceship).Methods("GET")
	r.HandleFunc("/spaceships", auth.AuthorizationMiddleware(controller.CreateSpaceship)).Methods("POST")
	r.HandleFunc("/spaceships/{id}", auth.AuthorizationMiddleware(controller.UpdateSpaceship)).Methods("PUT")
	r.HandleFunc("/spaceships/{id}", auth.AuthorizationMiddleware(controller.DeleteSpaceship)).Methods("DELETE")

	r.HandleFunc("/signup", controller.Signup).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")

	log.Println("listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
