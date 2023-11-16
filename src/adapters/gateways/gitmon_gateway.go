package gateways

import (
	"database/sql"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

type gitmonGateway struct {
	db *sql.DB
}

func NewGitmonGateway(db *sql.DB) *gitmonGateway {
	return &gitmonGateway{
		db: db,
	}
}

func (g *gitmonGateway) GetGitmonStatus(ownerID string) (types.GitmonStatus, error) {
	status, err := g.getGitmon(ownerID)
	if err != nil {
		return types.GitmonStatus{}, err
	}

	status.Skills, err = g.getSkills(status.GitmonID)
	if err != nil {
		return types.GitmonStatus{}, err
	}

	return status, nil
}

func (g *gitmonGateway) getGitmon(ownerID string) (types.GitmonStatus, error) {
	const query = `
		SELECT
			gitmon_id,
			gitmon_name,
			current_hp,
			current_attack,
			current_defence,
			current_speed
		FROM
			gitmons
		WHERE
			owner_id = $1
	`

	var gitmon types.GitmonStatus
	err := g.db.QueryRow(query, ownerID).Scan(
		&gitmon.GitmonID,
		&gitmon.Name,
		&gitmon.HP,
		&gitmon.Attack,
		&gitmon.Defence,
		&gitmon.Speed,
	)
	if err != nil {
		return types.GitmonStatus{}, err
	}

	return gitmon, nil
}

func (g *gitmonGateway) getSkills(gitmonID string) ([]types.Skill, error) {
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
		WHERE
			gitmon_id = $1
		AND 
			gitmon_skills.is_active = true
	`
	rows, err := g.db.Query(query, gitmonID)
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

	return skills, nil
}
