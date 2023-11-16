package interactors

import (
	"time"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/pkg/paseto"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/dai"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
)

type inGameInteractor struct {
	store      dai.InGameDai
	fbstore    dai.FirestoreDai
	wsstore    dai.WSDai
	maker      paseto.Maker
	outputport ports.InGameOutput
}

func NewIngameInteractor(store dai.InGameDai, fbstore dai.FirestoreDai, wsstore dai.WSDai, maker paseto.Maker, outputport ports.InGameOutput) ports.InGameInput {
	return &inGameInteractor{
		store:      store,
		fbstore:    fbstore,
		wsstore:    wsstore,
		maker:      maker,
		outputport: outputport,
	}
}

func (i *inGameInteractor) ReadyGame(reqBody input.ReadyGameRequest) (int, *output.ReadyGameResponse) {

	if len(reqBody.GameID) == 0 || len(reqBody.UserID) == 0 {
		// TODO: エラー処理
	}

	if err := i.store.ReadyGame(reqBody); err != nil {
		// TODO: エラー処理
	}

	ok, err := i.store.IsReady(reqBody.GameID)
	if err != nil {
		// TODO: エラー処理
	}

	if ok {
		token, err := i.maker.CreateToken(reqBody.GameID, config.Config.Server.AdminID, false, time.Hour*24*30*12*15)
		if err != nil {
			// TODO: エラー処理
		}
		go func() {
			if err := i.store.UpdateBlock(reqBody.GameID, true); err != nil {
				// TODO: エラー処理
			}

			if err := i.wsstore.Start(10, token); err != nil {
				// TODO: エラー処理
			}

			if err := i.store.UpdateBlock(reqBody.GameID, false); err != nil {
				// TODO: エラー処理
			}
		}()
	}

	return 0, nil
}

func (i *inGameInteractor) ActionGame(reqBody input.ActionGameRequest) (int, *output.ActionGameResponse) {
	if len(reqBody.GameID) == 0 {
		// TODO: エラー処理
	}
	ok, err := i.store.IsBlock(reqBody.GameID)
	if err != nil {
		// TODO: エラー処理
	}

	if !ok {
		// TODO: エラー処理
	}

	skill, err := i.store.GetSkill(reqBody.UserID, reqBody.CommandID)
	if err != nil {
		// TODO: エラー処理
	}

	if err := i.fbstore.CreateActionLog(skill); err != nil {

	}

	return 0, nil
}
