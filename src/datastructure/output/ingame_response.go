package output

import "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"

type ReadyGameResponse struct {
}

type ActionGameResponse struct {
	Skill types.Skill `json:"skill"`
}

type NextGameResponse struct {
}
