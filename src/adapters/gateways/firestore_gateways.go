package gateways

import (
	"context"
	"log"

	fb "firebase.google.com/go"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

type FireStoreGateway struct {
	app *fb.App
}

func NewFireStoreGateway(app *fb.App) *FireStoreGateway {
	return &FireStoreGateway{
		app: app,
	}
}

func (fs *FireStoreGateway) CreateGame(arg types.CreateGame) error {
	client, err := fs.app.Firestore(context.Background())
	if err != nil {
		log.Printf("Failed to create firestore client: %v", err)
		return err
	}

	gamedoc := client.Collection("games").Doc(arg.GameID)
	statuses := gamedoc.Collection("statuses")
	_, err = statuses.Doc(arg.OwnerID).Set(context.Background(), arg.OwnerGitmonStatus)

	return err
}

func (fs *FireStoreGateway) JoinGame(arg types.JoinGame) error {
	client, err := fs.app.Firestore(context.Background())
	if err != nil {
		log.Printf("Failed to create firestore client: %v", err)
		return err
	}

	gamedoc := client.Collection("games").Doc(arg.GameID)
	statuses := gamedoc.Collection("statuses")
	_, err = statuses.Doc(arg.UserID).Set(context.Background(), arg.UserIDGitmonStatus)

	return err
}
