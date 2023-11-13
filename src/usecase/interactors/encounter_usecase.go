package interactors

import (
	"time"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/dai"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EncounterInteracter struct {
	store      dai.EncounterDai
	outputPort ports.EncounterOutput
}

func NewEncounterInteracter(store dai.EncounterDai, outputPort ports.EncounterOutput) *EncounterInteracter {
	return &EncounterInteracter{
		store:      store,
		outputPort: outputPort,
	}
}

func (i *EncounterInteracter) Create(reqBody input.CreateEncounterReqeuest) (int, *output.CreateEncounterResponse) {
	if len(reqBody.UserID) == 0 || len(reqBody.EncountedUserID) == 0 {
		return i.outputPort.CreateEncounterResponse("", echo.ErrBadRequest)
	}
	return i.outputPort.CreateEncounterResponse(i.store.Create(types.CreateEncounter{
		EncounterID:     uuid.New().String(),
		UserID:          reqBody.UserID,
		EncountedUserID: reqBody.EncountedUserID,
		CreatedAt:       time.Now(),
	}))
}
