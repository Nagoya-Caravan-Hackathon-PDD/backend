package dai

import "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"

type EncounterDai interface {
	Create(arg types.CreateEncounter) error
}
