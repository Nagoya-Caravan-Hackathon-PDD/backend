package router

import (
	"net/http"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/controllers"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/gateways"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/persenters"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/interactors"
)

func (router *router) Health() {
	router.Mux.Handle("/v1/health", buildChain(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hc := controllers.NewHealthController(
				interactors.NewHealthInteractor(
					gateways.NewHealthGateway(router.db),
					persenters.NewHealthPresenter(w),
				),
			)

			switch r.Method {
			case http.MethodGet:
				hc.Health(w, r)
			default:
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		}),
		router.middleware.Recovery,
	))
}
