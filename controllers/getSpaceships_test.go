package controllers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"spacecraft/controllers"
	"spacecraft/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpaceShips(t *testing.T) {
	client := testutils.Setup()
	defer client.Close()

	ctx := context.Background()
	arm1 := client.Armament.Create().SetTitle("Ion Cannon").SetQty(10).SaveX(ctx)
	arm2 := client.Armament.Create().SetTitle("Laser Cannon").SetQty(20).SaveX(ctx)
	arm3 := client.Armament.Create().SetTitle("Missile Launcher").SetQty(30).SaveX(ctx)

	client.Spacecraft.Create().
		SetName("Star Destroyer").
		SetClass("Destroyer").
		SetCrew(5000).
		SetImage("https://starwars.com/star_destroyer.png").
		SetValue(10000).
		SetStatus("Destroyed").
		AddArmaments(arm1, arm2).
		SaveX(ctx)

	client.Spacecraft.Create().
		SetName("X-wing starfighter").
		SetClass("Starfighter").
		SetCrew(1).
		SetImage("https://starwars.com/x-wing.png").
		SetValue(2000).
		SetStatus("In crew training").
		AddArmaments(arm3).
		SaveX(ctx)

	ctrl := controllers.NewSpaceshipController(client)

	req, _ := http.NewRequest("GET", "/spaceships", nil)

	response := httptest.NewRecorder()

	ctrl.GetSpaceships(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	t.Log(response.Body.String())
}
