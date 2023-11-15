package controllers

import (
	"log"
	"net/http"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	_ "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
	"github.com/labstack/echo/v4"
)

type healthController struct {
	usecase ports.HealthInput
}

func NewHealthController(usecase ports.HealthInput) *healthController {
	return &healthController{usecase}
}

func (h *healthController) Health(ctx echo.Context) error {
	var reqQuery input.HealthRequest

	if err := ctx.Bind(&reqQuery); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(h.usecase.CheckDB(reqQuery))
}

func (h *healthController) AuthHealth(ctx echo.Context) error {
	var reqQuery input.HealthRequest

	if err := ctx.Bind(&reqQuery); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, err)
	}
	// payload := ctx.Get(middleware.PayloadContextKey).(*types.CustomClaims)
	// data, err := json.Marshal(payload)
	// if err != nil {
	// 	return ctx.JSON(http.StatusBadRequest, err)
	// }

	// log.Println(string(data))
	return ctx.JSON(h.usecase.CheckDB(reqQuery))
}
