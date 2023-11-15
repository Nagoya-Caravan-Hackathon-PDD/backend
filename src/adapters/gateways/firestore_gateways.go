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

const data = `
{
	"GameID":{
		"statuses":{
			"OwnerID":{
				// ここにステータス
			},
			"EnemyID":{
				// ここにステータス
			}
		},
		"action_logs":{
			1:{ // 1ターン目
				"OwnerID":{
					// ここにコマンドと結果
				},
				"EnemyID":{
					// ここにコマンドと結果
				}
			},
			2:{ // 1ターン目
				"OwnerID":{
					// ここにコマンドと結果
				},
				"EnemyID":{
					// ここにコマンドと結果
				}
			},
		}
	}
}
`

func (fs *FireStoreGateway) CreateGame(arg types.CreateGame) error {
	client, err := fs.app.Firestore(context.Background())
	if err != nil {
		log.Printf("Failed to create firestore client: %v", err)
		return err
	}

	gamedoc := client.Collection("games").Doc(arg.GameID)
	statuses := gamedoc.Collection("statuses")
	statuses.Doc(arg.OwnerID).Collection("ここに自分のギットモンステータスの構造体を入れる")
	gamedoc.Collection("action_logs")

	return nil
}
