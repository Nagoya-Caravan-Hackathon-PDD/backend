package controllers

import (
	"log"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
	"github.com/labstack/echo/v4"
)

type gameController struct {
	interactor ports.GameInput
}

func NewGameController(interactor ports.GameInput) *gameController {
	return &gameController{
		interactor: interactor,
	}
}

func (gc *gameController) Create(ctx echo.Context) error {
	var reqBody input.CreateGameRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(gc.interactor.CreateGame(reqBody))
}

func (gc *gameController) Join(ctx echo.Context) error {
	var reqBody input.JoinGameRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	log.Println(reqBody)
	return ctx.JSON(gc.interactor.JoinGame(reqBody))
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
