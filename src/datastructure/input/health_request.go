package input

type HealthRequest struct {
	CheckDB bool `query:"check_db"`
}
