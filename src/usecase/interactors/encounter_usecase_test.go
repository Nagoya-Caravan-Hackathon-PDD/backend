package interactors

import (
	"reflect"
	"testing"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/gateways"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/presenters"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

var users []types.CreateUser
var encounterID string

func TestCreateEncounter(t *testing.T) {
	eu := NewEncounterInteracter(gateways.NewEncounterGateway(dbconn), presenters.NewEncounterPresenter())
	mig.Up()
	defer migrateDown(t)

	for i := 0; i < 2; i++ {
		users = append(users, createUser(t))
	}
	testCases := []struct {
		name       string
		arg        input.CreateEncounterRequest
		wantStatus int
		wantOut    *output.CreateEncounterResponse
	}{
		{
			name: "success",
			arg: input.CreateEncounterRequest{
				UserID:          users[0].UserID,
				EncountedUserID: users[1].UserID,
			},
			wantStatus: 200,
			wantOut:    &output.CreateEncounterResponse{},
		},
		{
			name: "bad request uid empty",
			arg: input.CreateEncounterRequest{
				UserID:          "",
				EncountedUserID: users[1].UserID,
			},
			wantStatus: 400,
			wantOut:    nil,
		},
		{
			name: "bad request github_id empty",
			arg: input.CreateEncounterRequest{
				UserID:          users[0].UserID,
				EncountedUserID: "",
			},
			wantStatus: 400,
			wantOut:    nil,
		},
		{
			name: "duplicate encount",
			arg: input.CreateEncounterRequest{
				UserID:          users[0].UserID,
				EncountedUserID: users[1].UserID,
			},
			wantStatus: 409,
			wantOut:    nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, out := eu.Create(tc.arg)
			if status != tc.wantStatus {
				t.Errorf("want status %v, but got %v", tc.wantStatus, status)
			}
			if reflect.TypeOf(out) != reflect.TypeOf(tc.wantOut) {
				t.Errorf("want message %v, but got %v", tc.wantOut, out)
			}

			if out != nil {
				if out.EncounterID != "" {
					encounterID = out.EncounterID
				}
			}
		})
	}
	testListEncounter(t)
	testReadEncounter(t)
}

func testListEncounter(t *testing.T) {
	eu := NewEncounterInteracter(gateways.NewEncounterGateway(dbconn), presenters.NewEncounterPresenter())
	testCases := []struct {
		name       string
		arg        input.ListEncounterRequest
		wantStatus int
		wantOut    []*output.ListEncounterResponse
	}{
		{
			name: "success",
			arg: input.ListEncounterRequest{
				UserID: users[0].UserID,
			},
			wantStatus: 200,
			wantOut: []*output.ListEncounterResponse{
				{
					UserID:         users[0].UserID,
					EncoutedUserID: users[1].UserID,
				},
			},
		},
		{
			name: "bad request uid empty",
			arg: input.ListEncounterRequest{
				UserID: "",
			},
			wantStatus: 400,
			wantOut:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, out := eu.List(tc.arg)
			if status != tc.wantStatus {
				t.Errorf("got %v, want %v", status, tc.wantStatus)
			}
			if reflect.TypeOf(out) != reflect.TypeOf(tc.wantOut) {
				t.Errorf("got %v, want %v", out, tc.wantOut)
			}
		})
	}
}

func testReadEncounter(t *testing.T) {
	eu := NewEncounterInteracter(gateways.NewEncounterGateway(dbconn), presenters.NewEncounterPresenter())
	testCases := []struct {
		name       string
		arg        input.GetEncounterRequest
		wantStatus int
		wantOut    *output.ListEncounterResponse
	}{
		{
			name: "success",
			arg: input.GetEncounterRequest{
				EncounterID: encounterID,
			},
			wantStatus: 200,
			wantOut: &output.ListEncounterResponse{
				UserID:         users[0].UserID,
				EncoutedUserID: users[1].UserID,
			},
		},
		{
			name: "bad request uid empty",
			arg: input.GetEncounterRequest{
				EncounterID: "",
			},
			wantStatus: 400,
			wantOut:    nil,
		},
		{
			name: "not found",
			arg: input.GetEncounterRequest{
				EncounterID: "not_found",
			},
			wantStatus: 400,
			wantOut:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, out := eu.Read(tc.arg)
			if status != tc.wantStatus {
				t.Errorf("got %v, want %v", status, tc.wantStatus)
			}
			if reflect.TypeOf(out) != reflect.TypeOf(tc.wantOut) {
				t.Errorf("got %v, want %v", out, tc.wantOut)
			}
		})
	}
}
