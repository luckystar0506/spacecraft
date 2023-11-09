package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"spacecraft/ent"
	"spacecraft/ent/armament"
)

type SpaceshipController struct {
	client *ent.Client
}

type Armament struct {
	Title string `json:"title"`
	Qty   int    `json:"qty"`
}

type Spaceship struct {
	Name   string  `json:"name"`
	Class  string  `json:"class"`
	Crew   int     `json:"crew"`
	Image  string  `json:"image"`
	Value  float64 `json:"value"`
	Status string  `json:"status"`
}

type SpaceshipRequest struct {
	Spacecraft Spaceship  `json:"spacecraft"`
	Armaments  []Armament `json:"armaments"`
}

type Result struct {
	Success bool `json:"success"`
}

func NewSpaceshipController(client *ent.Client) *SpaceshipController {
	return &SpaceshipController{
		client: client,
	}
}

func (sc *SpaceshipController) CreateSpaceship(w http.ResponseWriter, r *http.Request) {
	var req SpaceshipRequest
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	tx, err := sc.client.Tx(ctx)

	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	newSpacecraft, err := tx.Spacecraft.Create().
		SetName(req.Spacecraft.Name).
		SetClass(req.Spacecraft.Class).
		SetCrew(req.Spacecraft.Crew).
		SetImage(req.Spacecraft.Image).
		SetValue(req.Spacecraft.Value).
		SetStatus(req.Spacecraft.Status).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to create spacecraft", http.StatusInternalServerError)
		return
	}

	armBuilders := make([]*ent.ArmamentCreate, len(req.Armaments))
	armTitles := make([]string, len(req.Armaments))
	for i, arm := range req.Armaments {
		armBuilders[i] = tx.Armament.Create().SetTitle(arm.Title)
		armTitles[i] = arm.Title
	}

	err = tx.Armament.CreateBulk(armBuilders...).OnConflict().DoNothing().Exec(ctx)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to create armaments", http.StatusInternalServerError)
		return
	}

	armaments, err := tx.Armament.Query().Where(armament.TitleIn(armTitles...)).All(ctx)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to query armaments", http.StatusInternalServerError)
		return
	}

	spacecraftarmamentBuilders := make([]*ent.SpacecraftArmamentCreate, len(req.Armaments))
	for i, arm := range req.Armaments {
		spacecraftarmamentBuilders[i] = tx.SpacecraftArmament.Create().SetSpacecraft(newSpacecraft).SetArmament(armaments[i]).SetQty(arm.Qty)
	}
	err = tx.SpacecraftArmament.CreateBulk(spacecraftarmamentBuilders...).Exec(ctx)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to create spacecraft armaments", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()

	if err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Result{Success: true})
}
