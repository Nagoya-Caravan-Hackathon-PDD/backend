package controllers

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	_ "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
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

// Game	godoc
//
// @Summary			Game Ready
// @Description	Game Ready
// @Tags				Game
// @Produce			json
// @param 			Authorization 		header 	string 							true 				"Authorization"
// @Param				ReadyGameRequest	body		input.ReadyGameRequest	true	"create game request"
// @Param				game_id						path		integer										true	"game id"
// @Success		200						{object}	nil			"success response"
// @Failure		400						{object}	nil														"error response"
// @Failure		500						{object}	nil														"error response"
// @Router		/game/{game_id}/ready				[POST]
func (gc *ingameController) Ready(ctx echo.Context) error {
	var reqBody input.ReadyGameRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(gc.interactor.ReadyGame(reqBody))
}

// Game	godoc
//
// @Summary			Game Action
// @Description	Game Action
// @Tags				Game
// @Produce			json
// @Param				game_id						path		integer										true	"game id"
// @param 			Authorization 		header 	string 							true 				"Authorization"
// @Param				ReadyGameRequest	body		input.ActionGameRequest	true	"create game request"
// @Success		200						{object}			nil			"success response"
// @Failure		400						{object}	nil														"error response"
// @Failure		500						{object}	nil														"error response"
// @Router		/game/{game_id}/action				[POST]
func (gc *ingameController) Action(ctx echo.Context) error {
	var reqBody input.ActionGameRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(gc.interactor.ActionGame(reqBody))
}
