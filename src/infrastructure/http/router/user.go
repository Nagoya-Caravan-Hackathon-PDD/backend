package router

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/controllers"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/gateways"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/presenters"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/interactors"
)

func (r *router) userRouter() {
	g := r.echo.Group("/v1")

	uc := controllers.NewUserController(
		interactors.NewUserInteractor(
			gateways.NewUserGateway(r.db),
			presenters.NewUserPresenter(),
		),
	)

	g.POST("/users", uc.CreateUser).Name = "CreateUser"
	g.GET("/users/:user_id", uc.GetUser).Name = "GetUsers"
	g.DELETE("/users/:user_id", uc.DeleteUsers).Name = "DeleteUsers"
}
