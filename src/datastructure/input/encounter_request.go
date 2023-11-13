package input

type CreateEncounterReqeuest struct {
	UserID          string `json:"user_id"`
	EncountedUserID string `json:"encounted_user_id"`
}
