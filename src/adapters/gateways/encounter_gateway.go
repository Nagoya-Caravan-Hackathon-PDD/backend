package gateways

import (
	"database/sql"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

type encounterGateway struct {
	db *sql.DB
}

func NewEncounterGateway(db *sql.DB) *encounterGateway {
	return &encounterGateway{
		db: db,
	}
}

func (eg *encounterGateway) getLocalEncounterCounts(userID, encountedUserID string) int {
	const query = `SELECT COUNT(*) FROM encounters WHERE from_user_id = $1 AND to_user_id = $2 AND created_at > NOW() - INTERVAL '1 hour'`
	rows := eg.db.QueryRow(query, userID, encountedUserID)

	var count int
	rows.Scan(&count)
	return count
}

func (eg *encounterGateway) Create(arg types.CreateEncounter) (string, error) {
	const query = `INSERT INTO encounters (encounter_id,from_user_id, to_user_id,created_at) VALUES ($1,$2,$3,$4)`
	count := eg.getLocalEncounterCounts(arg.UserID, arg.EncountedUserID)

	if count != 0 {
		return "", types.AlreadyExists
	}

	result, err := eg.db.Exec(query, arg.EncounterID, arg.UserID, arg.EncountedUserID, arg.CreatedAt)
	if err != nil {
		return "", err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if row != 1 {
		return "", types.InsertFailed
	}

	return arg.EncounterID, nil
}

func (eg *encounterGateway) ReadAll(arg input.ListEncounterRequest) ([]types.ReadEncounter, error) {
	const query = `SELECT * FROM encounters WHERE from_user_id = $1 LIMIT $2 OFFSET $3`
	row, err := eg.db.Query(query, arg.UserID, arg.PageSize, (arg.PageID-1)*arg.PageSize)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var encounters []types.ReadEncounter
	for row.Next() {
		var encounter types.ReadEncounter
		if err := row.Scan(&encounter.EncounterID, &encounter.UserID, &encounter.EncountedUserID, &encounter.CreatedAt); err != nil {
			return nil, err
		}
		encounters = append(encounters, encounter)
	}
	return encounters, nil
}

func (eg *encounterGateway) Read(encounterID string) (types.ReadEncounter, error) {
	const query = `SELECT * FROM encounters WHERE encounter_id = $1`
	row := eg.db.QueryRow(query, encounterID)

	var encounter types.ReadEncounter
	if err := row.Scan(&encounter.EncounterID, &encounter.UserID, &encounter.EncountedUserID, &encounter.CreatedAt); err != nil {
		return types.ReadEncounter{}, err
	}
	return encounter, nil
}
