package controllers

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/labstack/echo/v4"
)

type gameController struct{}

func NewGameController() *gameController {
	return &gameController{}
}

func (gc *gameController) Create(ctx echo.Context) error {
	var reqBody input.CreateGameRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return nil
}

func (gc *gameController) Join(ctx echo.Context) error {
	var reqBody input.JoinGameRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return nil
}

func (gc *gameController) List(ctx echo.Context) error {
	var reqBody input.ListGameRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return nil
}

func (gc *gameController) GetAny(ctx echo.Context) error {
	var reqBody input.GetAnyGameRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return nil
}
