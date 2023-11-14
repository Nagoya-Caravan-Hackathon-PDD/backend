package router

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/controllers"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/gateways"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/presenters"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/interactors"
)

func (r *router) encounterRoutes() {
	g := r.echo.Group("/v1")
	ec := controllers.NewEncounterController(
		interactors.NewEncounterInteracter(
			gateways.NewEncounterGateway(r.db),
			presenters.NewEncounterPresenter(),
		),
	)

	g.POST("/encounters", ec.Create)
	g.GET("/encounters", ec.List)
	g.GET("/encounters/:encounter_id", ec.Read)

}
