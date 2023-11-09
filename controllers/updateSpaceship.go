package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SpaceshipUpdateRequest struct {
	Spacecraft Spaceship  `json:"spacecraft"`
	Armaments  []Armament `json:"armaments"`
}

// 1. Update the Spacecraft fields (like name, class, crew, image, value, status)
//
// 2. For each Armament in the request:
//
// - If a SpacecraftArmament association already exists for this Armament and Spacecraft, just update that record with the new quantity.
//
// - If an Armament already exists but there's no SpacecraftArmament association, create the SpacecraftArmament association.
//
// - If the Armament doesn't exist yet, create the Armament, then create the SpacecraftArmament.
func (sc *SpaceshipController) UpdateSpaceship(w http.ResponseWriter, r *http.Request) {
	var req SpaceshipUpdateRequest
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	ctx := context.Background()

	tx, err := sc.client.Tx(ctx)
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	spacecraft, err := tx.Spacecraft.
		UpdateOneID(id).
		SetName(req.Spacecraft.Name).
		SetClass(req.Spacecraft.Class).
		SetCrew(req.Spacecraft.Crew).
		SetImage(req.Spacecraft.Image).
		SetValue(req.Spacecraft.Value).
		SetStatus(req.Spacecraft.Status).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to update spacecraft", http.StatusInternalServerError)
		return
	}

	for _, armamentReq := range req.Armaments {
		armamentID, err := tx.Armament.Create().SetTitle(armamentReq.Title).OnConflict().DoNothing().ID(ctx)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to create armament", http.StatusInternalServerError)
			return
		}

		err = tx.SpacecraftArmament.Create().SetArmamentID(armamentID).SetSpacecraft(spacecraft).SetQty(armamentReq.Qty).OnConflict().UpdateQty().Exec(ctx)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to update spacecraftarmament", http.StatusInternalServerError)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Result{Success: true})
}
