package router

import "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/controllers"

func (r *router) encounterRoutes() {
	g := r.echo.Group("/v1")
	ec := controllers.NewEncounterController()

	g.POST("/encounters", ec.Create)
}
