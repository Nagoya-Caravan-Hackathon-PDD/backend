package gateways

import (
	"testing"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

func TestCreate(t *testing.T) {
	defer migrateDown(t)
	g := NewUserGateway(dbconn)

	testCases := []struct {
		name    string
		arg     types.CreateUser
		wantErr error
	}{
		{
			name: "success",
			arg: types.CreateUser{
				UserID:   "test_user_id",
				GitHubID: "test_github_id",
			},
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := g.Create(tc.arg); err != tc.wantErr {
				t.Errorf("want: %v, got: %v", tc.wantErr, err)
			}
		})
	}

	testRead(t)
	testDelete(t)
}

func testRead(t *testing.T) {

	g := NewUserGateway(dbconn)

	testCases := []struct {
		name    string
		arg     string
		wantErr error
	}{
		{
			name:    "success",
			arg:     "test_user_id",
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := g.Read(tc.arg); err != tc.wantErr {
				t.Errorf("want: %v, got: %v", tc.wantErr, err)
			}
		})
	}
}

func testDelete(t *testing.T) {

	g := NewUserGateway(dbconn)

	testCases := []struct {
		name    string
		arg     string
		wantErr error
	}{
		{
			name:    "success",
			arg:     "test_user_id",
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := g.Delete(tc.arg); err != tc.wantErr {
				t.Errorf("want: %v, got: %v", tc.wantErr, err)
			}
		})
	}
}
