package ports

import "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"

type HealthInput interface {
	CheckDB(input.HealthRequest) error
}

type HealthOutput interface {
	Success() error
	Failed(error) error
}
