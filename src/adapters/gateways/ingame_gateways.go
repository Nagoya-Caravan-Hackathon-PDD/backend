package gateways

import (
	"database/sql"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

type inGameGateway struct {
	db *sql.DB
}

func NewIngameGateway(db *sql.DB) *inGameGateway {
	return &inGameGateway{
		db: db,
	}
}

func (g *inGameGateway) GetGameInfo(gameid string) (*types.GameTemp, error) {
	const query = `
		SELECT
			game_id,
			owner_id,
			owner_is_ok,
			enemy_id,
			enemy_is_ok,
			is_block
		FROM
			game_temp
		WHERE
			game_id = $1
	`
	var gametemp types.GameTemp
	if err := g.db.QueryRow(query, gameid).Scan(&gametemp.GameID, &gametemp.OwnerID, &gametemp.OwnerIsOK, &gametemp.EnemyID, &gametemp.EnemyIsOK, &gametemp.IsBlock); err != nil {
		return nil, err
	}
	return &gametemp, nil
}

func (g *inGameGateway) ReadyGame(arg input.ReadyGameRequest) error {
	info, err := g.GetGameInfo(arg.GameID)
	if err != nil {
		return err
	}

	var query = `
		UPDATE
			game_temp
		SET
	`
	if info.OwnerID == arg.UserID {
		query += `
			owner_is_ok = true
		`
	} else {
		query += `
			enemy_is_ok = true
		`
	}

	query += `
		WHERE
			game_id = $1
	`

	if _, err := g.db.Exec(query, arg.GameID); err != nil {
		return err
	}

	return nil
}

func (g *inGameGateway) IsReady(gameid string) (bool, error) {
	const query = `
		SELECT
			game_id,
			owner_is_ok,
			enemy_is_ok,
			is_block
		FROM
			game_temp
		WHERE 
			game_id = $1
	`
	var gametemp types.GameTemp

	if err := g.db.QueryRow(query, gameid).Scan(&gametemp.GameID, &gametemp.OwnerIsOK, &gametemp.EnemyIsOK, &gametemp.IsBlock); err != nil {
		return false, err
	}

	return gametemp.OwnerIsOK && gametemp.EnemyIsOK, nil
}

func (g *inGameGateway) UpdateBlock(gameID string, status bool) error {
	const query = `
		UPDATE
			game_temp
		SET
			is_block = $1
		WHERE
			game_id = $2
	`
	if _, err := g.db.Exec(query, status, gameID); err != nil {
		return err
	}
	return nil
}

func (g *inGameGateway) IsBlock(gameID string) (bool, error) {
	const query = `
		SELECT
			game_id,
			is_block
		FROM
			game_temp
		WHERE
			game_id = $1
	`
	var gametemp types.GameTemp

	if err := g.db.QueryRow(query, gameID).Scan(&gametemp.GameID, &gametemp.IsBlock); err != nil {
		return false, err
	}

	return gametemp.IsBlock, nil
}
