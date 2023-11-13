package presenters

import (
	"net/http"
	"reflect"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type encounterPresenter struct {
}

func NewEncounterPresenter() *encounterPresenter {
	return &encounterPresenter{}
}

func (p *encounterPresenter) CreateEncounterResponse(err error) (int, *output.CreateEncounterResponse) {
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

	return http.StatusOK, &output.CreateEncounterResponse{}
}
