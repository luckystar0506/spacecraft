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
	armamentBulk := client.Armament.CreateBulk(
		client.Armament.Create().SetTitle("Ion Cannon"),
		client.Armament.Create().SetTitle("Laser Cannon"),
		client.Armament.Create().SetTitle("Missile Launcher"),
	).SaveX(ctx)

	spacecraftBulk := client.Spacecraft.CreateBulk(
		client.Spacecraft.Create().
			SetName("Star Destroyer").
			SetClass("Destroyer").
			SetCrew(5000).
			SetImage("https://starwars.com/star_destroyer.png").
			SetValue(10000).
			SetStatus("Destroyed"),
		client.Spacecraft.Create().
			SetName("X-wing starfighter").
			SetClass("Starfighter").
			SetCrew(1).
			SetImage("https://starwars.com/x-wing.png").
			SetValue(2000).
			SetStatus("In crew training"),
	).SaveX(ctx)

	_ = client.SpacecraftArmament.CreateBulk(
		client.SpacecraftArmament.Create().SetArmament(armamentBulk[0]).SetSpacecraft(spacecraftBulk[0]).SetQty(10),
		client.SpacecraftArmament.Create().SetArmament(armamentBulk[1]).SetSpacecraft(spacecraftBulk[0]).SetQty(20),
		client.SpacecraftArmament.Create().SetArmament(armamentBulk[1]).SetSpacecraft(spacecraftBulk[1]).SetQty(30),
		client.SpacecraftArmament.Create().SetArmament(armamentBulk[2]).SetSpacecraft(spacecraftBulk[1]).SetQty(40),
	).SaveX(ctx)

	ctrl := controllers.NewSpaceshipController(client)

	req, _ := http.NewRequest("GET", "/spaceships", nil)

	response := httptest.NewRecorder()

	ctrl.GetSpaceships(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	t.Log(response.Body.String())
}
