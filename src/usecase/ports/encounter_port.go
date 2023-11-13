package ports

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

type EncounterInput interface {
	Create(arg input.CreateEncounterReqeuest) (int, *output.CreateEncounterResponse)
	List(args input.ListEncounterRequest) (int, []*output.ListEncounterResponse)
}

type EncounterOutput interface {
	CreateEncounterResponse(encounterID string, err error) (int, *output.CreateEncounterResponse)
	GetEncounterResponse(args []types.ReadEncounter, err error) (int, []*output.ListEncounterResponse)
}
