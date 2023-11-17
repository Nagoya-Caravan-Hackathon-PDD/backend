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

func (g *inGameGateway) CreateGame(gameID, ownerID string, speed int) error {
	const query = `
		INSERT INTO
			game_temp
		(
			game_id,
			owner_id,
			owner_is_ok,
			enemy_id,
			enemy_is_ok,
			is_block,
			owner_speed,
			is_first,
			game_index,
			is_end
		)
		VALUES
		(
			$1,
			$2,
			false,
			'',
			false,
			false,
			$3,
			true,
			0,
			false
		)
	`
	if _, err := g.db.Exec(query, gameID, ownerID, speed); err != nil {
		return err
	}
	return nil
}

func (g *inGameGateway) JoinGame(gameID, userID string, speed int) error {
	gameinfo, err := g.GetGameInfo(gameID)
	if err != nil {
		return err
	}
	var query = `
	UPDATE
		game_temp
	SET
		enemy_id = $1,
	`

	if gameinfo.OwnerSpeed < speed {
		query += `
			is_first = false
		`
	} else {
		query += `
			is_first = true
		`
	}

	query += `
	WHERE
		game_id = $2
	`
	if _, err := g.db.Exec(query, userID, gameID); err != nil {
		return err
	}
	return nil
}

func (g *inGameGateway) GetGameInfo(gameid string) (*types.GameTemp, error) {
	const query = `
		SELECT
			game_id,
			owner_id,
			owner_is_ok,
			enemy_id,
			enemy_is_ok,
			is_block,
			owner_speed,
			is_first,
			game_index
		FROM
			game_temp
		WHERE
			game_id = $1
	`
	var gametemp types.GameTemp
	if err := g.db.QueryRow(query, gameid).Scan(&gametemp.GameID, &gametemp.OwnerID, &gametemp.OwnerIsOK, &gametemp.EnemyID, &gametemp.EnemyIsOK, &gametemp.IsBlock, &gametemp.OwnerSpeed, &gametemp.IsFirst, &gametemp.GameIndex); err != nil {
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

func (g *inGameGateway) UpdateBlock(gameID string, status bool) (int, error) {
	var query = `
		UPDATE
			game_temp
		SET
			is_block = $1
	`
	// ロック解除したらindexを増やす
	if status {
		query += `
			,game_index = game_index + 1
		`
	}
	query += `
		WHERE
			game_id = $2
		RETURNING game_index
	`

	var index int
	if err := g.db.QueryRow(query, status, gameID).Scan(&index); err != nil {
		return 0, err
	}

	return index, nil
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

func (g *inGameGateway) GetSkill(userID string, skillID int) (*types.Skill, error) {
	const query = `
		SELECT
			skills.skill_id,
			skills.skill_name,
			skills.description,
			skills.skilltype,
			skills.value
		FROM
			gitmon_skills
		LEFT OUTER JOIN
			skills
		ON
			gitmon_skills.skill_id = skills.skill_id
		LEFT OUTER JOIN
			gitmons
		ON
			gitmon_skills.gitmon_id = gitmons.gitmon_id
		WHERE
			gitmons.owner_id = $1
		AND
			skills.skill_id = $2	
		AND 
			gitmon_skills.is_active = true
	`

	rows, err := g.db.Query(query, userID, skillID)
	if err != nil {
		return nil, err
	}

	var skills []types.Skill
	for rows.Next() {
		var skill types.Skill
		err := rows.Scan(
			&skill.SkillID,
			&skill.SkillName,
			&skill.Description,
			&skill.SkillType,
			&skill.SkillValue,
		)
		if err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	if len(skills) == 0 {
		return nil, types.ErrDontHaveSkill
	}

	return &skills[0], nil
}

func (i *inGameGateway) ReasetReady(gameID string) error {
	const query = `
		UPDATE
			game_temp
		SET
			owner_is_ok = false,
			enemy_is_ok = false
		WHERE
			game_id = $1
	`
	if _, err := i.db.Exec(query, gameID); err != nil {
		return err
	}
	return nil
}

func (i *inGameGateway) EndGame(gameID string) error {
	const query = `
	UPDATE
		game_temp
	SET
		is_end = true
	WHERE
		game_id = $1
	`
	if _, err := i.db.Exec(query, gameID); err != nil {
		return err
	}
	return nil
}

func (i *inGameGateway) IsEnd(gameID string) (bool, error) {
	const query = `
		SELECT
			is_end
		FROM
			game_temp
		WHERE
			game_id = $1
	`
	var isEnd bool
	if err := i.db.QueryRow(query, gameID).Scan(&isEnd); err != nil {
		return false, err
	}
	return isEnd, nil
}
