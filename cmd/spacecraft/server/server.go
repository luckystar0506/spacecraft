package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"spacecraft/controllers"
	"spacecraft/ent"
	"spacecraft/routers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     string
)

func init() {
	log.Println("Initializing server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser = getEnvVariable("DATABASE_USER")
	dbPassword = getEnvVariable("DATABASE_PASSWORD")
	dbName = getEnvVariable("DATABASE_NAME")
	dbHost = getEnvVariable("DATABASE_HOST")
	dbPort = getEnvVariable("DATABASE_PORT")
}

func RunServer() {
	log.Println("Running server...")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	client, err := ent.Open("mysql", dbURL)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	controller := controllers.NewSpaceshipController(client)

	r := mux.NewRouter()

	spaceshipRouter := r.PathPrefix("/spaceships").Subrouter()
	routers.SpaceshipRoutes(spaceshipRouter, controller)

	authRouter := r.PathPrefix("/auth").Subrouter()
	routers.AuthRoutes(authRouter, controller)

	log.Println("listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getEnvVariable(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("%s not found in .env file", key)
	}
	return value
}
