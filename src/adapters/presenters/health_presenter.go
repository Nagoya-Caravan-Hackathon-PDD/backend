package presenters

import (
	"net/http"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
)

type healthPresenter struct {
}

func NewHealthPresenter() *healthPresenter {
	return &healthPresenter{}
}

func (p *healthPresenter) Health(err error) (int, output.HealthResponse) {
	if err != nil {
		return http.StatusInternalServerError, output.HealthResponse{Message: "NG"}
	}
	return http.StatusOK, output.HealthResponse{Message: "OK"}
}
