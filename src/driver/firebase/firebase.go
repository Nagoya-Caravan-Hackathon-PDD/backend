package firebase

import (
	"context"

	fb "firebase.google.com/go"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/cmd/config"
	"google.golang.org/api/option"
)

func FbApp(path string) (*fb.App, error) {
	firebaseconfig := &fb.Config{
		StorageBucket: config.Config.Firebase.StorageBucket,
	}

	serviceAccount := option.WithCredentialsFile(path)
	app, err := fb.NewApp(context.Background(), firebaseconfig, serviceAccount)
	if err != nil {
		return nil, err
	}

	return app, nil
}
