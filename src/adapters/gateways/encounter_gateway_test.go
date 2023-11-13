package gateways

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/pkg/utils"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
	"github.com/jackc/pgx/v5/pgconn"
)

func createEncounter(t *testing.T, users []types.CreateUser) types.CreateEncounter {
	arg := types.CreateEncounter{
		EncounterID:     utils.RandomString(10),
		UserID:          users[1].UserID,
		EncountedUserID: users[0].UserID,
		CreatedAt:       time.Now(),
	}
	if err := NewEncounterGateway(dbconn).Create(arg); err != nil {
		t.Fatal(err)
	}
	return arg
}
func TestCreateEncounter(t *testing.T) {
	migrateUp(t)
	defer migrateDown(t)

	eg := NewEncounterGateway(dbconn)
	var users []types.CreateUser
	for i := 0; i < 2; i++ {
		users = append(users, createOneUser(t))
	}

	createEncounter(t, users)

	testCases := []struct {
		name        string
		arg         types.CreateEncounter
		wantErr     error
		wantErrCode string
	}{
		{
			name: "success",
			arg: types.CreateEncounter{
				EncounterID:     "encounter_id",
				UserID:          users[0].UserID,
				EncountedUserID: users[1].UserID,
				CreatedAt:       time.Now(),
			},
			wantErr: nil,
		},
		{
			name: "already exists",
			arg: types.CreateEncounter{
				EncounterID:     "encounter_id1",
				UserID:          users[1].UserID,
				EncountedUserID: users[0].UserID,
				CreatedAt:       time.Now().Add(time.Second),
			},
			wantErr: types.AlreadyExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := eg.Create(tc.arg)
			if reflect.TypeOf(err) != reflect.TypeOf(tc.wantErr) {
				t.Errorf("got %v, want %s", err, reflect.TypeOf(tc.wantErr))
			} else {
				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) {
					if pgErr.Code != tc.wantErrCode {
						t.Errorf("got %v, want %v", pgErr.Code, tc.wantErrCode)
					}
				}
			}
		})
	}
}
