package presenters

import (
	"log"
	"net/http"
	"reflect"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/labstack/echo/v4"
)

type GamePresenter struct {
}

func NewGamePresenter() *GamePresenter {
	return &GamePresenter{}
}

func (gp *GamePresenter) CreateGame(token, gameID string, err error) (int, *output.CreateGameResponse) {
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

	return http.StatusOK, &output.CreateGameResponse{
		GameServerToken: token,
		GameID:          gameID,
	}
}
func (gp *GamePresenter) JoinGame(token, gameID string, err error) (int, *output.JoinGameResponse) {
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

	return http.StatusOK, &output.JoinGameResponse{
		GameServerToken: token,
		GameID:          gameID,
	}
}
func (gp *GamePresenter) ListGame([]*input.GetAnyGameRequest, error) (int, []*input.GetAnyGameRequest) {
	return http.StatusOK, nil
}
func (gp *GamePresenter) GetAnyGame(output.GetAnyGameResponse, error) (int, *output.GetAnyGameResponse) {
	return http.StatusOK, nil
}
