package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"spacecraft/ent/spacecraft"
)

type SpacecraftResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (sc *SpaceshipController) GetSpaceships(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	spacecraftQuery := sc.client.Spacecraft.Query().Where(spacecraft.DeletedAtIsNil())

	if name := r.URL.Query().Get("name"); name != "" {
		spacecraftQuery.Where(spacecraft.NameEQ(name))
	}

	if class := r.URL.Query().Get("class"); class != "" {
		spacecraftQuery.Where(spacecraft.ClassEQ(class))
	}

	if status := r.URL.Query().Get("status"); status != "" {
		spacecraftQuery.Where(spacecraft.StatusEQ(status))
	}

	dbSpacecrafts, err := spacecraftQuery.All(ctx)

	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get spacecrafts", http.StatusInternalServerError)
		return
	}

	spacecraftsRes := make([]SpacecraftResponse, len(dbSpacecrafts))

	for i, dbSpacecraft := range dbSpacecrafts {
		spacecraftsRes[i] = SpacecraftResponse{
			ID:     dbSpacecraft.ID,
			Name:   dbSpacecraft.Name,
			Status: dbSpacecraft.Status,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spacecraftsRes)
}
