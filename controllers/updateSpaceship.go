package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SpaceshipUpdateRequest struct {
	Spacecraft Spaceship `json:"spacecraft"`
}

func (sc *SpaceshipController) UpdateSpaceship(w http.ResponseWriter, r *http.Request) {
	var req SpaceshipUpdateRequest
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	ctx := context.Background()

	_, err = sc.client.Spacecraft.
		UpdateOneID(id).
		SetName(req.Spacecraft.Name).
		SetClass(req.Spacecraft.Class).
		SetCrew(req.Spacecraft.Crew).
		SetImage(req.Spacecraft.Image).
		SetValue(req.Spacecraft.Value).
		SetStatus(req.Spacecraft.Status).
		Save(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Result{Success: true})
}
