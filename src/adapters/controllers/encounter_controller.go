package controllers

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	_ "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
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

// Encounter	godoc
//
// @Summary		Create Encounter
// @Description	Create Encount entory
// @Tags		Encounter
// @Produce		json
// @Param		CreateEncounterRequest	body		input.CreateEncounterRequest	true	"create encounter request"
// @Success		200						{object}	output.CreateEncounterResponse			"success response"
// @Failure		400						{object}	nil										"error response"
// @Failure		500						{object}	nil										"error response"
// @Router		/encounters				[POST]
func (ec *encounterController) Create(ctx echo.Context) error {
	var reqBody input.CreateEncounterRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(ec.interactor.Create(reqBody))
}

// Encounter	godoc
//
// @Summary		Get All Encounters
// @Description	Get All Encounters
// @Tags		Encounter
// @Produce		json
// @Param		ListEncounterRequest	query		input.ListEncounterRequest	true	"list encounter request"
// @Success		200						{array}		output.ListEncounterResponse		"success response"
// @Failure		400						{object}	nil									"error response"
// @Failure		500						{object}	nil									"error response"
// @Router		/encounters				[GET]
func (ec *encounterController) List(ctx echo.Context) error {
	var reqBody input.ListEncounterRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(ec.interactor.List(reqBody))
}

// Encounter	godoc
//
// @Summary		Get All Encounters
// @Description	Get All Encounters
// @Tags		Encounter
// @Produce		json
// @Param		encounter_id			path		string						true	"list encounter request"
// @Success		200						{object}	output.ListEncounterResponse		"success response"
// @Failure		400						{object}	nil									"error response"
// @Failure		500						{object}	nil									"error response"
// @Router		/encounters/:encounter_id			[GET]
func (ec *encounterController) Read(ctx echo.Context) error {
	var reqBody input.GetEncounterRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(ec.interactor.Read(reqBody))
}
