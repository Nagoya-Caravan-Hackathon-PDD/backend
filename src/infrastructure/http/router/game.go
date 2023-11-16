package router

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/pkg/paseto"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/controllers"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/gateways"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/presenters"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/interactors"
)

func (r *router) gameRouter() {
	maker, err := paseto.NewPasetoMaker(config.Config.Paseto.SecretKey)
	if err != nil {
		panic(err)
	}
	gc := controllers.NewGameController(
		interactors.NewGameInteractor(
			gateways.NewFireStoreGateway(r.app),
			gateways.NewGitmonGateway(r.db),
			maker,
			presenters.NewGamePresenter(),
		),
	)

	g := r.echo.Group("/v1")
	g.POST("/game", gc.Create)
	g.POST("/game/:game_id", gc.Join)
	g.GET("/game", gc.List)
	g.GET("/game/:game_id", gc.GetAny)
}
