package interactors

import (
	"fmt"
	"log"
	"time"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/pkg/calc"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/pkg/paseto"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/dai"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
	"github.com/labstack/echo/v4"
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
	log.Println(reqBody)
	if len(reqBody.GameID) == 0 || len(reqBody.UserID) == 0 {
		return i.outputport.ReadyGame(echo.ErrBadRequest)
	}

	if err := i.store.ReadyGame(reqBody); err != nil {
		return i.outputport.ReadyGame(err)
	}

	ok, err := i.store.IsReady(reqBody.GameID)
	if err != nil {
		return i.outputport.ReadyGame(err)
	}

	if ok {
		token, err := i.maker.CreateToken(reqBody.GameID, config.Config.Server.AdminID, false, time.Hour*24*30*12*15)
		if err != nil {
			return i.outputport.ReadyGame(err)
		}
		go func() {
			ind, err := i.store.UpdateBlock(reqBody.GameID, true)
			if err != nil {
				log.Println(err)
				return
			}
			if ind == 1 {
				if err := i.wsstore.Start(4, token); err != nil {
					log.Println(err)
					return
				}
			}

			if err := i.wsstore.Turn(10, token); err != nil {
				log.Println(err)
				return
			}

			if _, err := i.store.UpdateBlock(reqBody.GameID, false); err != nil {
				log.Println(err)
				return
			}

			// 計算処理
			if err := i.calc(reqBody.GameID); err != nil {
				log.Println(err)
				return
			}
			// Wsに通知
		}()
	}

	return i.outputport.ReadyGame(nil)
}

func (i *inGameInteractor) calc(gameID string) error {
	token, err := i.maker.CreateToken(gameID, config.Config.Server.AdminID, false, time.Hour*24*30*12*15)
	if err != nil {
		log.Println(err)
	}
	temps, err := i.store.GetGameInfo(gameID)
	if err != nil {
		log.Println(err)
	}

	var (
		infos  [2]*types.GitmonStatus
		skills [2]*types.Skill
	)

	infos[0], err = i.fbstore.GetGitmonStatus(gameID, temps.OwnerID)
	if err != nil {
		log.Println(err)
	}

	infos[1], err = i.fbstore.GetGitmonStatus(gameID, temps.EnemyID.String)
	if err != nil {
		log.Println(err)
	}

	var (
		owner = -1
		enemy = -1
	)

	if temps.IsFirst {
		owner = 0
		enemy = 1
	} else {
		owner = 1
		enemy = 0
	}

	if temps.IsFirst {
		skills[owner], err = i.fbstore.GetActionLog(fmt.Sprintf("%d", temps.GameIndex), gameID, temps.OwnerID)
		if err != nil {
			log.Println(err)
		}
		skills[enemy], err = i.fbstore.GetActionLog(fmt.Sprintf("%d", temps.GameIndex), gameID, temps.EnemyID.String)
		if err != nil {
			log.Println(err)
		}
	} else {
		skills[owner], err = i.fbstore.GetActionLog(fmt.Sprintf("%d", temps.GameIndex), gameID, temps.OwnerID)
		if err != nil {
			log.Println(err)
		}
		skills[enemy], err = i.fbstore.GetActionLog(fmt.Sprintf("%d", temps.GameIndex), gameID, temps.EnemyID.String)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println(infos[0].Name, infos[1].Name)
	// 先行の計算
	switch skills[0].SkillType {
	case types.TypeAttack:
		dmg := infos[1].HP - calc.Attack(infos[0].Attack*int(skills[0].SkillValue), infos[1].Defence)
		infos[1].HP = dmg
		log.Println("first attack", dmg)
	case types.TypeDefence:
		infos[0].Defence += infos[0].Defence * int(skills[0].SkillValue)
		// 一旦考慮しない
	}

	// 後攻の計算
	switch skills[1].SkillType {
	case types.TypeAttack:
		dmg := infos[0].HP - calc.Attack(infos[1].Attack*int(skills[1].SkillValue), infos[0].Defence)
		infos[0].HP = dmg
		log.Println("first attack", dmg)
	case types.TypeDefence:
		infos[1].Defence += infos[1].Defence * int(skills[1].SkillValue)
		// 一旦考慮しない
	}

	log.Println(&infos)
	// 結果をFBに反映

	if err := i.fbstore.JoinGame(types.JoinGame{
		GameID:             gameID,
		UserID:             temps.OwnerID,
		UserIDGitmonStatus: *infos[owner],
	}); err != nil {
		return err
	}

	if err := i.fbstore.JoinGame(types.JoinGame{
		GameID:             gameID,
		UserID:             temps.EnemyID.String,
		UserIDGitmonStatus: *infos[enemy],
	}); err != nil {
		return err
	}

	if err := i.wsstore.Result(0, token); err != nil {
		log.Println(err)
	}

	if infos[0].HP <= 0 || infos[1].HP <= 0 {
		// WSに通知
		if err := i.wsstore.End(0, token); err != nil {
			log.Println(err)
		}
		// DBに通知
		if err := i.store.EndGame(gameID); err != nil {
			log.Println(err)
		}
	}

	if err := i.store.ReasetReady(gameID); err != nil {
		return err
	}
	return nil
}

func (i *inGameInteractor) ActionGame(reqBody input.ActionGameRequest) (int, *output.ActionGameResponse) {
	if len(reqBody.GameID) == 0 {
		return i.outputport.ActionGame(nil, echo.ErrBadRequest)
	}
	info, err := i.store.GetGameInfo(reqBody.GameID)
	if err != nil {
		return i.outputport.ActionGame(nil, err)
	}

	if !info.IsBlock {
		return i.outputport.ActionGame(nil, echo.ErrBadRequest)
	}

	skill, err := i.store.GetSkill(reqBody.UserID, reqBody.CommandID)
	if err != nil {
		return i.outputport.ActionGame(nil, err)
	}

	if err := i.fbstore.AddActionLog(fmt.Sprintf("%d", info.GameIndex), reqBody.GameID, reqBody.UserID, *skill); err != nil {
		return i.outputport.ActionGame(nil, err)
	}

	return i.outputport.ActionGame(skill, nil)
}
