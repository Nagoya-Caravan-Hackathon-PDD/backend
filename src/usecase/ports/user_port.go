package ports

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

type UserInput interface {
	Create(types.CreateUser) (int, *output.CreateUserResponse)
	Read(input.GetUser) (int, *output.ReadUserResponse)
	Delete(input.DeleteUsers) (int, *output.DeleteUserResponse)
}

type UserOutput interface {
	Create(error) (int, *output.CreateUserResponse)
	Read(error) (int, *output.ReadUserResponse)
	Delete(error) (int, *output.DeleteUserResponse)
}
