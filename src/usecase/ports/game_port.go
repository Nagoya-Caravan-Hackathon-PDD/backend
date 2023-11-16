package ports

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
)

type GameInput interface {
	CreateGame(input.CreateGameRequest) (int, *output.CreateGameResponse)
	JoinGame(input.JoinGameRequest) (int, *output.JoinGameResponse)
	ListGame(input.ListGameRequest) (int, []*input.GetAnyGameRequest)
	GetAnyGame(input.GetAnyGameRequest) (int, *output.GetAnyGameResponse)
}

type GameOutput interface {
	CreateGame(token, gameID string, err error) (int, *output.CreateGameResponse)
	JoinGame(token string, err error) (int, *output.JoinGameResponse)
	ListGame([]*input.GetAnyGameRequest, error) (int, []*input.GetAnyGameRequest)
	GetAnyGame(output.GetAnyGameResponse, error) (int, *output.GetAnyGameResponse)
}
