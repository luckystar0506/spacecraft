package main

import (
	"context"
	"log"
	"net/http"
	"spacecraft/controllers"
	"spacecraft/ent"
	"spacecraft/middleware/auth"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "user:password@tcp(127.0.0.1:3306)/spacecraft")
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
