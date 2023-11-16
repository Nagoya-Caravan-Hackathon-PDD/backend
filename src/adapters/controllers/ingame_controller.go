package controllers

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
	"github.com/labstack/echo/v4"
)

type ingameController struct {
	interactor ports.InGameInput
}

func NewInGameController(interactor ports.InGameInput) *ingameController {
	return &ingameController{
		interactor: interactor,
	}
}

func (gc *gameController) Ready(ctx echo.Context) error {
	var reqBody input.ReadyGameRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return nil
}

func (gc *gameController) Action(ctx echo.Context) error {
	var reqBody input.ActionGameRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return nil
}

func (gc *gameController) FinTurn(ctx echo.Context) error {
	var reqBody input.FinTurnRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return nil
}
