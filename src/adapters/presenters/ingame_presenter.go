package presenters

import (
	"log"
	"net/http"
	"reflect"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
	"github.com/labstack/echo/v4"
)

type ingamePresenter struct {
}

func NewIngamePresenter() *ingamePresenter {
	return &ingamePresenter{}
}

func (p *ingamePresenter) ReadyGame(err error) (int, *output.ReadyGameResponse) {
	if err != nil {
		log.Println(err)
		switch reflect.TypeOf(err) {
		case reflect.TypeOf(echo.ErrBadRequest):
			return http.StatusBadRequest, nil
		case reflect.TypeOf(echo.ErrInternalServerError):
			return http.StatusInternalServerError, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	return http.StatusOK, &output.ReadyGameResponse{}
}

func (p *ingamePresenter) ActionGame(skill *types.Skill, err error) (int, *output.ActionGameResponse) {
	if err != nil {
		switch reflect.TypeOf(err) {
		case reflect.TypeOf(echo.ErrBadRequest):
			return http.StatusBadRequest, nil
		case reflect.TypeOf(types.ErrDontHaveSkill):
			return http.StatusBadRequest, nil
		case reflect.TypeOf(echo.ErrInternalServerError):
			return http.StatusInternalServerError, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	return http.StatusOK, &output.ActionGameResponse{
		Skill: *skill,
	}
}
