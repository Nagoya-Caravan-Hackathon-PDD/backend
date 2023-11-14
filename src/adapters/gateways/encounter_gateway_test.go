package gateways

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/pkg/utils"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lib/pq"
)

var users []types.CreateUser

func createEncounter(t *testing.T) types.CreateEncounter {
	arg := types.CreateEncounter{
		EncounterID:     utils.RandomString(10),
		UserID:          users[1].UserID,
		EncountedUserID: users[0].UserID,
		CreatedAt:       time.Now(),
	}
	if _, err := NewEncounterGateway(dbconn).Create(arg); err != nil {
		t.Fatal(err)
	}
	return arg
}
func TestCreateEncounter(t *testing.T) {
	migrateUp(t)
	defer migrateDown(t)

	eg := NewEncounterGateway(dbconn)

	for i := 0; i < 2; i++ {
		users = append(users, createOneUser(t))
	}

	createEncounter(t)

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
		{
			name: "user not found",
			arg: types.CreateEncounter{
				EncounterID:     "encounter_id2",
				UserID:          "not_found_user_id",
				EncountedUserID: users[0].UserID,
				CreatedAt:       time.Now().Add(time.Second * 2),
			},
			wantErr:     &pq.Error{},
			wantErrCode: pgerrcode.ForeignKeyViolation,
		},
		{
			name: "encounted user not found",
			arg: types.CreateEncounter{
				EncounterID:     "encounter_id3",
				UserID:          users[0].UserID,
				EncountedUserID: "not_found_user_id",
				CreatedAt:       time.Now().Add(time.Second * 3),
			},
			wantErr:     &pq.Error{},
			wantErrCode: pgerrcode.ForeignKeyViolation,
		},
		{
			name: "user and encounted user are same",
			arg: types.CreateEncounter{
				EncounterID:     "encounter_id4",
				UserID:          users[0].UserID,
				EncountedUserID: users[0].UserID,
				CreatedAt:       time.Now().Add(time.Second * 4),
			},
			wantErr:     &pq.Error{},
			wantErrCode: pgerrcode.CheckViolation,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := eg.Create(tc.arg)
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

	testReadOne(t)
	testReadAll(t)
}

func testReadOne(t *testing.T) {
	eg := NewEncounterGateway(dbconn)

	testCases := []struct {
		name    string
		arg     string
		want    types.ReadEncounter
		wantErr error
	}{
		{
			name: "success",
			arg:  "encounter_id",
			want: types.ReadEncounter{
				EncounterID:     "encounter_id",
				UserID:          users[1].UserID,
				EncountedUserID: users[0].UserID,
				CreatedAt:       time.Now().Truncate(time.Second).UTC(),
			},
			wantErr: nil,
		},
		{
			name:    "not found",
			arg:     "not_found_encounter_id",
			want:    types.ReadEncounter{},
			wantErr: sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := eg.Read(tc.arg)
			if err != tc.wantErr {
				t.Errorf("got %v, want %v", err, tc.wantErr)
			}
			got.CreatedAt = got.CreatedAt.Truncate(time.Second)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func testReadAll(t *testing.T) {
	eg := NewEncounterGateway(dbconn)

	testCases := []struct {
		name    string
		arg     input.ListEncounterRequest
		wantLen int
		wantErr error
	}{
		{
			name: "success",
			arg: input.ListEncounterRequest{
				UserID:   users[1].UserID,
				PageID:   1,
				PageSize: 10,
			},
			wantLen: 1,
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := eg.ReadAll(tc.arg)
			if err != tc.wantErr {
				t.Errorf("got %v, want %v", err, tc.wantErr)
			}

			if len(got) != tc.wantLen {
				t.Errorf("got %v, want %v", len(got), tc.wantLen)
			}
		})
	}
}
