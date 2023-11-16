package interactors

import (
	"time"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/pkg/paseto"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/dai"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type gameInteractor struct {
	fbstore     dai.FirestoreDai
	gitmonstore dai.GitmonDai
	maker       paseto.Maker
	outputport  ports.GameOutput
}

func NewGameInteractor(fbstore dai.FirestoreDai, gitmonstore dai.GitmonDai, maker paseto.Maker, outputport ports.GameOutput) *gameInteractor {
	return &gameInteractor{
		fbstore:     fbstore,
		gitmonstore: gitmonstore,
		outputport:  outputport,
		maker:       maker,
	}
}

func (gi *gameInteractor) CreateGame(arg input.CreateGameRequest) (int, *output.CreateGameResponse) {
	if len(arg.OwnerID) == 0 {
		return gi.outputport.CreateGame("", "", echo.ErrBadRequest)
	}

	gameID := uuid.New().String()

	token, err := gi.maker.CreateToken(gameID, arg.OwnerID, true, time.Duration(30*time.Minute))
	if err != nil {
		return gi.outputport.CreateGame("", "", err)
	}
	gitmon, err := gi.gitmonstore.GetGitmonStatus(arg.OwnerID)
	if err != nil {
		return gi.outputport.CreateGame("", "", err)
	}

	if err := gi.fbstore.CreateGame(types.CreateGame{
		GameID:            gameID,
		OwnerID:           arg.OwnerID,
		CreatedAt:         time.Now(),
		OwnerGitmonStatus: gitmon,
	}); err != nil {
		return gi.outputport.CreateGame("", "", err)
	}

	return gi.outputport.CreateGame(token, gameID, nil)
}

func (gi *gameInteractor) JoinGame(arg input.JoinGameRequest) (int, *output.JoinGameResponse) {
	if len(arg.GameID) == 0 || len(arg.UserID) == 0 {
		return gi.outputport.JoinGame("", "", echo.ErrBadRequest)
	}

	token, err := gi.maker.CreateToken(arg.GameID, arg.UserID, false, time.Duration(30*time.Minute))
	if err != nil {
		return gi.outputport.JoinGame("", "", err)
	}
	gitmon, err := gi.gitmonstore.GetGitmonStatus(arg.UserID)
	if err != nil {
		return gi.outputport.JoinGame("", "", err)
	}

	if err := gi.fbstore.JoinGame(types.JoinGame{
		GameID:             arg.GameID,
		UserID:             arg.UserID,
		UserIDGitmonStatus: gitmon,
	}); err != nil {
		return gi.outputport.JoinGame("", "", err)
	}

	return gi.outputport.JoinGame(token, arg.GameID, nil)
}

func (gi *gameInteractor) ListGame(arg input.ListGameRequest) (int, []*input.GetAnyGameRequest) {
	if len(arg.UserID) == 0 {
		return gi.outputport.ListGame(nil, echo.ErrBadRequest)
	}

	if arg.PageID == 0 {
		arg.PageID = 1
	}

	if arg.PageSize == 0 {
		arg.PageSize = 10
	}

	return gi.outputport.ListGame(nil, nil)
}

func (gi *gameInteractor) GetAnyGame(arg input.GetAnyGameRequest) (int, *output.GetAnyGameResponse) {
	if len(arg.GameID) == 0 {
		return gi.outputport.GetAnyGame(output.GetAnyGameResponse{}, echo.ErrBadRequest)
	}

	return gi.outputport.GetAnyGame(output.GetAnyGameResponse{}, nil)
}
