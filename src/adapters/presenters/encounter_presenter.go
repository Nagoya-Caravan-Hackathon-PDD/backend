package presenters

import (
	"net/http"
	"reflect"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type encounterPresenter struct {
}

func NewEncounterPresenter() *encounterPresenter {
	return &encounterPresenter{}
}

func (p *encounterPresenter) CreateEncounterResponse(encounterID string, err error) (int, *output.CreateEncounterResponse) {
	if err != nil {
		switch reflect.TypeOf(err) {
		case reflect.TypeOf(&pq.Error{}):
			return http.StatusBadRequest, &output.CreateEncounterResponse{}
		case reflect.TypeOf(echo.ErrBadRequest):
			return http.StatusBadRequest, &output.CreateEncounterResponse{}
		case reflect.TypeOf(echo.ErrInternalServerError):
			return http.StatusInternalServerError, &output.CreateEncounterResponse{}
		default:
			return http.StatusInternalServerError, &output.CreateEncounterResponse{}
		}
	}

	return http.StatusOK, &output.CreateEncounterResponse{
		EncounterID: encounterID,
	}
}

func (p *encounterPresenter) ListEncounterResponse(args []types.ReadEncounter, err error) (int, []*output.ListEncounterResponse) {
	if err != nil {
		switch reflect.TypeOf(err) {
		case reflect.TypeOf(&pq.Error{}):
			return http.StatusBadRequest, nil
		case reflect.TypeOf(echo.ErrBadRequest):
			return http.StatusBadRequest, nil
		case reflect.TypeOf(echo.ErrInternalServerError):
			return http.StatusInternalServerError, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	var res []*output.ListEncounterResponse
	for _, arg := range args {
		res = append(res, &output.ListEncounterResponse{
			EncounterID:    arg.EncounterID,
			UserID:         arg.UserID,
			EncoutedUserID: arg.EncountedUserID,
			CreatedAt:      arg.CreatedAt,
		})
	}
	return http.StatusOK, res
}

func (p *encounterPresenter) GetEncounterResponse(args types.ReadEncounter, err error) (int, *output.ListEncounterResponse) {
	if err != nil {
		switch reflect.TypeOf(err) {
		case reflect.TypeOf(&pq.Error{}):
			return http.StatusBadRequest, nil
		case reflect.TypeOf(echo.ErrBadRequest):
			return http.StatusBadRequest, nil
		case reflect.TypeOf(echo.ErrInternalServerError):
			return http.StatusInternalServerError, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	return http.StatusOK, &output.ListEncounterResponse{
		EncounterID:    args.EncounterID,
		UserID:         args.UserID,
		EncoutedUserID: args.EncountedUserID,
		CreatedAt:      args.CreatedAt,
	}
}
