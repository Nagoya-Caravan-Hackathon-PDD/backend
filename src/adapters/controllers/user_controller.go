package controllers

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type userController struct {
	interactor ports.UserInput
}

func NewUserController(interactor ports.UserInput) *userController {
	return &userController{
		interactor: interactor,
	}
}

func (uc *userController) CreateUser(ctx echo.Context) error {
	var reqBody input.CreateUser

	if err := ctx.Bind(&reqBody); err != nil {
		return echo.ErrBadRequest
	}
	uid := uuid.New().String()
	return ctx.JSON(uc.interactor.Create(types.CreateUser{
		UserID:   uid,
		GitHubID: reqBody.GitHub,
	}))
}

func (uc *userController) GetUser(ctx echo.Context) error {
	var reqQuery input.GetUser

	if err := ctx.Bind(&reqQuery); err != nil {
		return echo.ErrBadRequest
	}
	return nil
}

func (uc *userController) DeleteUsers(ctx echo.Context) error {
	var reqQuery input.DeleteUsers

	if err := ctx.Bind(&reqQuery); err != nil {
		return echo.ErrBadRequest
	}
	return nil
}
