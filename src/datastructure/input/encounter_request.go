package input

type CreateEncounterReqeuest struct {
	UserID          string `json:"user_id"`
	EncountedUserID string `json:"encounted_user_id"`
}

type GetEncounterRequest struct {
	UserID   string `param:"user_id"`
	PageSize int    `query:"page_size"`
	PageID   int    `query:"page_id"`
}
