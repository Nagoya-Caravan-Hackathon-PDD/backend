package router

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/controllers"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/gateways"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/presenters"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/interactors"
)

func (router *router) Health() {
	hc := controllers.NewHealthController(
		interactors.NewHealthInteractor(
			gateways.NewHealthGateway(router.db),
			presenters.NewHealthPresenter(),
		),
	)

	router.echo.GET("/health", hc.Health)
}
