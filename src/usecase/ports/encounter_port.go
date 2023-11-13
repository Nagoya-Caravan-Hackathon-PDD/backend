package ports

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
)

type EncounterInput interface {
	Create(arg input.CreateEncounterReqeuest) (int, *output.CreateEncounterResponse)
}

type EncounterOutput interface {
	CreateEncounterResponse(encounterID string, err error) (int, *output.CreateEncounterResponse)
}
