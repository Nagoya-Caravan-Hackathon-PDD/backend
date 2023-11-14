package dai

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

type EncounterDai interface {
	Create(arg types.CreateEncounter) (string, error)
	ReadAll(arg input.ListEncounterRequest) ([]types.ReadEncounter, error)
	Read(encounterID string) (types.ReadEncounter, error)
}
