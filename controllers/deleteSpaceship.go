package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (sc *SpaceshipController) DeleteSpaceship(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	ctx := context.Background()

	err := sc.client.Spacecraft.UpdateOneID(id).SetDeletedAt(time.Now()).Exec(ctx)

	if err != nil {
		http.Error(w, "Failed to delete spacecraft", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Result{Success: true})
}
