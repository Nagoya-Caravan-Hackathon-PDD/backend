package interactors

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/dai"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
)

type HealthInteractor struct {
	store      dai.HealthDai
	outputport ports.HealthOutput
}

func NewHealthInteractor(store dai.HealthDai, outputport ports.HealthOutput) *HealthInteractor {
	return &HealthInteractor{store, outputport}
}

func (i *HealthInteractor) CheckDB(reqQuery input.HealthRequest) error {
	if reqQuery.CheckDB {
		if err := i.store.Ping(); err != nil {
			return i.outputport.Failed(err)
		}
	}

	return i.outputport.Success()
}
