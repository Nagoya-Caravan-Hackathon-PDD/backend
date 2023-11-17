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

func (fs *FireStoreGateway) AddActionLog(index, gameID, userID string, skill types.Skill) error {
	client, err := fs.app.Firestore(context.Background())
	if err != nil {
		log.Printf("Failed to create firestore client: %v", err)
		return err
	}

	accesslog := client.Collection("games").Doc(gameID).Collection("log")
	_, err = accesslog.Doc("action").Collection(index).Doc(userID).Set(context.Background(), skill)

	return err
}

func (fs *FireStoreGateway) GetActionLog(index, gameID, userID string) (*types.Skill, error) {
	client, err := fs.app.Firestore(context.Background())
	if err != nil {
		log.Printf("Failed to create firestore client: %v", err)
		return nil, err
	}
	var skill *types.Skill

	accesslog := client.Collection("games").Doc(gameID).Collection("log")
	dsnap, err := accesslog.Doc("action").Collection(index).Doc(userID).Get(context.Background())
	if err != nil {
		return nil, err
	}

	if err := dsnap.DataTo(&skill); err != nil {
		return nil, err
	}

	return skill, nil
}

func (fs *FireStoreGateway) GetGitmonStatus(gameID, userID string) (*types.GitmonStatus, error) {
	client, err := fs.app.Firestore(context.Background())
	if err != nil {
		log.Printf("Failed to create firestore client: %v", err)
		return nil, err
	}
	var gitmon *types.GitmonStatus

	dsnap, err := client.Collection("games").Doc(gameID).Collection("statuses").Doc(userID).Get(context.Background())

	if err != nil {
		return nil, err
	}

	if err := dsnap.DataTo(&gitmon); err != nil {
		return nil, err
	}

	return gitmon, nil
}
