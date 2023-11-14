package output

import "time"

type CreateEncounterResponse struct {
	EncounterID string `json:"encounter_id"`
}
type ListEncounterResponse struct {
	EncounterID    string    `json:"encounter_id"`
	UserID         string    `json:"user_id"`
	EncoutedUserID string    `json:"encouted_user_id"`
	CreatedAt      time.Time `json:"created_at"`
}
