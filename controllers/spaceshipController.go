package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"spacecraft/ent"
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
	Success bool `json:"spacecraft"`
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	arms := make([]*ent.Armament, len(req.Armaments))

	for i, arm := range req.Armaments {
		arms[i], err = tx.Armament.Create().SetTitle(arm.Title).SetQty(arm.Qty).Save(ctx)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	_, err = tx.Spacecraft.Create().
		SetName(req.Spacecraft.Name).
		SetClass(req.Spacecraft.Class).
		SetCrew(req.Spacecraft.Crew).
		SetImage(req.Spacecraft.Image).
		SetValue(req.Spacecraft.Value).
		SetStatus(req.Spacecraft.Status).
		AddArmaments(arms...).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tx.Commit()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Result{Success: true})
}
