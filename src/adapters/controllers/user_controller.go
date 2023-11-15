package controllers

import (
	"log"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	_ "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
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

// User	godoc
//
// @Summary		Create User
// @Description Create User
// @Tags		User
// @Produce		json
// @Param		CreateUser	body		input.CreateUser			true	"create user request"
// @Success		200						{object}	output.CreateUserResponse			"success response"
// @Failure		400						{object}	nil									"error response"
// @Failure		409						{object}	nil									"error response"
// @Failure		500						{object}	nil									"error response"
// @Router		/users					[POST]
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

// User	godoc
//
// @Summary		Get User
// @Description Get any User
// @Tags		User
// @Produce		json
// @Param			user_id				path			string						true		"create user request"
// @Success		200						{object}	output.ReadUserResponse		"success response"
// @Failure		400						{object}	nil												"error response"
// @Failure		409						{object}	nil												"error response"
// @Failure		500						{object}	nil												"error response"
// @Router		/users/{user_id}			[GET]
func (uc *userController) GetUser(ctx echo.Context) error {
	var reqQuery input.GetUser
	if err := ctx.Bind(&reqQuery); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(uc.interactor.Read(reqQuery))
}

// User	godoc
//
// @Summary		Delete User
// @Description Delete User
// @Tags		User
// @Produce		json
// @Param		user_id					path		string						true	"create user request"
// @Success		200						{object}	output.DeleteUserResponse			"success response"
// @Failure		400						{object}	nil									"error response"
// @Failure		409						{object}	nil									"error response"
// @Failure		500						{object}	nil									"error response"
// @Router		/users/{user_id}			[DELETE]
func (uc *userController) DeleteUsers(ctx echo.Context) error {
	var reqQuery input.DeleteUsers

	if err := ctx.Bind(&reqQuery); err != nil {
		return echo.ErrBadRequest
	}
	log.Println(reqQuery)
	return nil
}
