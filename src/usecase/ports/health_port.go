package ports

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
)

type HealthInput interface {
	CheckDB(input.HealthRequest) (int, output.HealthResponse)
}

type HealthOutput interface {
	Health(error) (int, output.HealthResponse)
}
