package controllers

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
	"github.com/labstack/echo/v4"
)

type encounterController struct {
	interactor ports.EncounterInput
}

func NewEncounterController(interactor ports.EncounterInput) *encounterController {
	return &encounterController{
		interactor: interactor,
	}
}

func (ec *encounterController) Create(ctx echo.Context) error {
	var reqBody input.CreateEncounterReqeuest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(ec.interactor.Create(reqBody))
}

func (ec *encounterController) List(ctx echo.Context) error {
	var reqBody input.ListEncounterRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(ec.interactor.List(reqBody))
}
