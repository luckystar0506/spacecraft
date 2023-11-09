package routers

import (
	"net/http"
	"spacecraft/controllers"
	"spacecraft/middleware/auth"

	"github.com/gorilla/mux"
)

func SpaceshipRoutes(r *mux.Router, controller *controllers.SpaceshipController) {
	r.HandleFunc("", controller.GetSpaceships).Methods(http.MethodGet)
	r.HandleFunc("/{id}", controller.GetSpaceship).Methods(http.MethodGet)
	r.HandleFunc("", auth.AuthorizationMiddleware(controller.CreateSpaceship)).Methods(http.MethodPost)
	r.HandleFunc("/{id}", auth.AuthorizationMiddleware(controller.UpdateSpaceship)).Methods(http.MethodPut)
	r.HandleFunc("/{id}", auth.AuthorizationMiddleware(controller.DeleteSpaceship)).Methods(http.MethodDelete)
}
