package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"spacecraft/ent"
	"spacecraft/ent/spacecraft"
	"strconv"

	"github.com/gorilla/mux"
)

type SpacecraftDetailResponse struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Class     string     `json:"class"`
	Crew      int        `json:"crew"`
	Image     string     `json:"image"`
	Value     float64    `json:"value"`
	Status    string     `json:"status"`
	Armaments []Armament `json:"armaments"`
}

func (sc *SpaceshipController) GetSpaceship(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	dbSpacecraft, err := sc.client.Spacecraft.
		Query().
		Where(
			spacecraft.IDEQ(id),
			spacecraft.DeletedAtIsNil(),
		).
		WithArmaments(func(saq *ent.SpacecraftArmamentQuery) {
			saq.WithArmament()
		}).
		Only(ctx)

	if err != nil {
		// check if it's notfound error
		if ent.IsNotFound(err) {
			http.Error(w, "Spacecraft not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, "Failed to get spacecraft", http.StatusInternalServerError)
			return
		}
	}

	armaments := make([]Armament, len(dbSpacecraft.Edges.Armaments))
	for i, arm := range dbSpacecraft.Edges.Armaments {
		armaments[i] = Armament{
			Title: arm.Edges.Armament.Title,
			Qty:   arm.Qty,
		}
	}

	spacecraftRes := SpacecraftDetailResponse{
		ID:        dbSpacecraft.ID,
		Name:      dbSpacecraft.Name,
		Class:     dbSpacecraft.Class,
		Crew:      dbSpacecraft.Crew,
		Image:     dbSpacecraft.Image,
		Value:     dbSpacecraft.Value,
		Status:    dbSpacecraft.Status,
		Armaments: armaments,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spacecraftRes)
}
