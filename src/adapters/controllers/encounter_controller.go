package controllers

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	_ "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/infrastructure/http/middleware"
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
// @param 	Authorization 					header 	string 												true 	"Authorization"
// @Param		CreateEncounterRequest	body		input.CreateEncounterRequest	true	"create encounter request"
// @Success		200						{object}	output.CreateEncounterResponse			"success response"
// @Failure		400						{object}	nil										"error response"
// @Failure		500						{object}	nil										"error response"
// @Router		/v1/encounters				[POST]
func (ec *encounterController) Create(ctx echo.Context) error {
	var reqBody input.CreateEncounterRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}

	payload := ctx.Get(middleware.PayloadContextKey).(*types.CustomClaims)
	if reqBody.UserID != payload.UserId {
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
// @param 		Authorization 				header 	string 											true 	"Authorization"
// @Param			ListEncounterRequest	query		input.ListEncounterRequest	true	"list encounter request"
// @Success		200						{array}		output.ListEncounterResponse		"success response"
// @Failure		400						{object}	nil									"error response"
// @Failure		500						{object}	nil									"error response"
// @Router		/v1/encounters		[GET]
func (ec *encounterController) List(ctx echo.Context) error {
	var reqQuery input.ListEncounterRequest
	if err := ctx.Bind(&reqQuery); err != nil {
		return echo.ErrBadRequest
	}

	payload := ctx.Get(middleware.PayloadContextKey).(*types.CustomClaims)
	if reqQuery.UserID != payload.UserId {
		return echo.ErrBadRequest
	}

	return ctx.JSON(ec.interactor.List(reqQuery))
}

// Encounter	godoc
//
// @Summary		Get All Encounters
// @Description	Get All Encounters
// @Tags		Encounter
// @Produce		json
// @param 	Authorization 		header 			string 										true 	"Authorization"
// @Param		encounter_id			path				string										true	"list encounter request"
// @Success		200						{object}			output.ListEncounterResponse		"success response"
// @Failure		400						{object}			nil															"error response"
// @Failure		500						{object}			nil															"error response"
// @Router		/v1/encounters/{encounter_id}	[GET]
func (ec *encounterController) Read(ctx echo.Context) error {
	var reqQuery input.GetEncounterRequest
	if err := ctx.Bind(&reqQuery); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(ec.interactor.Read(reqQuery))
}
