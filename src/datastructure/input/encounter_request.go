package input

type CreateEncounterRequest struct {
	UserID          string `json:"user_id"`
	EncountedUserID string `json:"encounted_user_id"`
}
type ListEncounterRequest struct {
	UserID   string `query:"user_id"`
	PageSize uint   `query:"page_size"`
	PageID   uint   `query:"page_id"`
}

type GetEncounterRequest struct {
	EncounterID string `param:"encounter_id"`
}
