package router

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/pkg/paseto"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/controllers"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/gateways"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/gateways/ws"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/presenters"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/interactors"
)

func (r *router) ingameRouter() {
	maker, err := paseto.NewPasetoMaker(config.Config.Paseto.SecretKey)
	if err != nil {
		panic(err)
	}

	gc := controllers.NewInGameController(
		interactors.NewIngameInteractor(
			gateways.NewIngameGateway(r.db),
			gateways.NewFireStoreGateway(r.app),
			ws.NewWSRequest(),
			maker,
			presenters.NewIngamePresenter(),
		),
	)

	g := r.echo.Group("/v1")

	g.POST("/game/:game_id/ready", gc.Ready)
	g.POST("/game/:game_id/action", gc.Action)
	g.POST("/game/:game_id/fin_turn", gc.FinTurn)
}
