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

func (i *UserInteractor) Read(input.GetUser) (int, *output.ReadUserResponse) {
	return i.outputport.Read(nil)
}
func (i *UserInteractor) Delete(input.DeleteUsers) (int, *output.DeleteUserResponse) {
	return i.outputport.Delete(nil)
}
