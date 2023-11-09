package routers

import (
	"net/http"
	"spacecraft/controllers"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, controller *controllers.SpaceshipController) {
	r.HandleFunc("/signup", controller.Signup).Methods(http.MethodPost)
	r.HandleFunc("/login", controller.Login).Methods(http.MethodPost)
}
