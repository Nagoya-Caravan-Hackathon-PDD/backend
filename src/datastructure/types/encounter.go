package types

import "time"

type CreateEncounter struct {
	EncounterID     string
	UserID          string
	EncountedUserID string
	CreatedAt       time.Time
}
