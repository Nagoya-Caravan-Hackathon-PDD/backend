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

func (i *EncounterInteracter) Create(reqBody input.CreateEncounterRequest) (int, *output.CreateEncounterResponse) {
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

func (i *EncounterInteracter) List(arg input.ListEncounterRequest) (int, []*output.ListEncounterResponse) {

	if len(arg.UserID) == 0 {
		return i.outputPort.ListEncounterResponse(nil, echo.ErrBadRequest)
	}
	if arg.PageID == 0 {
		arg.PageID = 1
	}
	if arg.PageSize == 0 {
		arg.PageSize = 10
	}
	return i.outputPort.ListEncounterResponse(i.store.ReadAll(arg))
}

func (i *EncounterInteracter) Read(arg input.GetEncounterRequest) (int, *output.ListEncounterResponse) {
	if len(arg.EncounterID) == 0 {
		return i.outputPort.GetEncounterResponse(types.ReadEncounter{}, echo.ErrBadRequest)
	}
	return i.outputPort.GetEncounterResponse(i.store.Read(arg.EncounterID))
}
