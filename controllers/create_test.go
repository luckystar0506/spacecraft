package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"spacecraft/controllers"
	"spacecraft/testutils"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCreateSpaceship(t *testing.T) {
	client := testutils.Setup()
	defer client.Close()

	ctrl := controllers.NewSpaceshipController(client)

	spaceshipRequest := controllers.SpaceshipRequest{
		Spacecraft: controllers.Spaceship{
			Name:   "Starship",
			Class:  "Spacefighter",
			Crew:   3500,
			Image:  "http://xyz.com/abc.png",
			Value:  2345.5,
			Status: "operational",
		},
		Armaments: []controllers.Armament{
			{
				Title: "Turbo Laser",
				Qty:   60,
			},
			{
				Title: "Ion Cannons",
				Qty:   60,
			},
			{
				Title: "Tractor Beam",
				Qty:   10,
			},
		},
	}

	jsonReq, _ := json.Marshal(spaceshipRequest)
	req, _ := http.NewRequest("POST", "/spaceships", bytes.NewBuffer(jsonReq))
	response := httptest.NewRecorder()

	ctrl.CreateSpaceship(response, req)

	assert.Equal(t, 200, response.Code)

	var result controllers.Result
	_ = json.Unmarshal(response.Body.Bytes(), &result)

	assert.Equal(t, true, result.Success)
}
