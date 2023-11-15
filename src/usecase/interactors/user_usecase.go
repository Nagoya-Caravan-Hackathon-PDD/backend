package interactors

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/dai"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
	"github.com/labstack/echo/v4"
)

type UserInteractor struct {
	store      dai.UserDai
	outputport ports.UserOutput
}

func NewUserInteractor(store dai.UserDai, outputport ports.UserOutput) *UserInteractor {
	return &UserInteractor{store, outputport}
}

func (i *UserInteractor) Create(reqBody types.CreateUser) (int, *output.CreateUserResponse) {
	if len(reqBody.GitHubID) == 0 || len(reqBody.UserID) == 0 {
		return i.outputport.Create(echo.ErrBadRequest)
	}
	return i.outputport.Create(i.store.Create(reqBody))
}

func (i *UserInteractor) Read(arg input.GetUser) (int, *output.ReadUserResponse) {
	if len(arg.UserID) == 0 {
		return i.outputport.Read(nil, echo.ErrBadRequest)
	}
	return i.outputport.Read(i.store.Read(arg.UserID))
}

func (i *UserInteractor) Delete(reqQuery input.DeleteUsers) (int, *output.DeleteUserResponse) {
	if len(reqQuery.UserID) == 0 {
		return i.outputport.Delete(echo.ErrBadRequest)
	}
	return i.outputport.Delete(i.store.Delete(reqQuery.UserID))
}
