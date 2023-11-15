package dai

import "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"

type UserDai interface {
	Create(arg types.CreateUser) (err error)
	Read(userID string) (*types.User, error)
	Delete(userID string) (err error)
}
