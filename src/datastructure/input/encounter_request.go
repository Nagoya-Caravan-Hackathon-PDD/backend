package input

type CreateEncounterReqeuest struct {
	UserID          string `json:"user_id"`
	EncountedUserID string `json:"encounted_user_id"`
}
type ListEncounterRequest struct {
	UserID   string `query:"user_id"`
	PageSize uint   `query:"page_size"`
	PageID   uint   `query:"page_id"`
}
