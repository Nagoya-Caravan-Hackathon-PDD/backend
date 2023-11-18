package controllers

import (
	"log"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	_ "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
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

// Game	godoc
//
// @Summary			Create Game
// @Description	Create Game
// @Tags				Game
// @Produce		json
// @param 	Authorization 		header 	string 						true 				"Authorization"
// @Param		CreateGameRequest	body		input.CreateGameRequest	true	"create game request"
// @Success		200						{object}	output.CreateGameResponse			"success response"
// @Failure		400						{object}	nil														"error response"
// @Failure		500						{object}	nil														"error response"
// @Router		/v1/game				[POST]
func (gc *gameController) Create(ctx echo.Context) error {
	var reqBody input.CreateGameRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(gc.interactor.CreateGame(reqBody))
}

// Game	godoc
//
// @Summary			Join Game
// @Description	Join Game
// @Tags				Game
// @Produce			json
// @param 			Authorization 		header 	string 							true 				"Authorization"
// @Param				game_id						path		string										true	"game id"
// @Param				JoinGameRequest	body		input.JoinGameRequest	true	"create game request"
// @Success		200						{object}	output.JoinGameResponse			"success response"
// @Failure		400						{object}	nil														"error response"
// @Failure		500						{object}	nil														"error response"
// @Router		/v1/game/{game_id}				[POST]
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
