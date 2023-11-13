package interactors

import (
	"net/http"
	"testing"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/gateways"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/presenters"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

func TestCreate(t *testing.T) {
	defer migrateDown(t)
	ui := NewUserInteractor(gateways.NewUserGateway(dbconn), presenters.NewUserPresenter())
	testCases := []struct {
		name       string
		arg        types.CreateUser
		wantStatus int
		wantBody   *output.CreateUserResponse
	}{
		{
			name: "success",
			arg: types.CreateUser{
				UserID:   "test_user_id",
				GitHubID: "test_github_id",
			},
			wantStatus: http.StatusOK,
			wantBody: &output.CreateUserResponse{
				Message: "create successful",
			},
		},
		{
			name: "bad request uid empty",
			arg: types.CreateUser{
				UserID:   "",
				GitHubID: "test_github_id",
			},
			wantStatus: http.StatusBadRequest,
			wantBody: &output.CreateUserResponse{
				Message: "Bad Request",
			},
		},
		{
			name: "bad request github_id empty",
			arg: types.CreateUser{
				UserID:   "",
				GitHubID: "test_github_id",
			},
			wantStatus: http.StatusBadRequest,
			wantBody: &output.CreateUserResponse{
				Message: "Bad Request",
			},
		},
		{
			name: "conflict",
			arg: types.CreateUser{
				UserID:   "test_user_id",
				GitHubID: "test_github_id",
			},
			wantStatus: http.StatusConflict,
			wantBody: &output.CreateUserResponse{
				Message: "Conflict",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, body := ui.Create(tc.arg)
			if status != tc.wantStatus {
				t.Errorf("got %v, want %v", status, tc.wantStatus)
			}
			if body.Message != tc.wantBody.Message {
				t.Errorf("got %v, want %v", body.Message, tc.wantBody.Message)
			}
		})
	}
	testDelete(t)
}

func testDelete(t *testing.T) {
	ui := NewUserInteractor(gateways.NewUserGateway(dbconn), presenters.NewUserPresenter())
	testCases := []struct {
		name       string
		arg        input.DeleteUsers
		wantStatus int
		wantBody   *output.DeleteUserResponse
	}{
		{
			name: "success",
			arg: input.DeleteUsers{
				UserID: "test_user_id",
			},
			wantStatus: http.StatusOK,
			wantBody: &output.DeleteUserResponse{
				Message: "delete successful",
			},
		},
		{
			name: "bad request uid empty",
			arg: input.DeleteUsers{
				UserID: "",
			},
			wantStatus: http.StatusBadRequest,
			wantBody: &output.DeleteUserResponse{
				Message: "Bad Request",
			},
		},
		{
			name: "bad request uid not found",
			arg: input.DeleteUsers{
				UserID: "test_user_id",
			},
			wantStatus: http.StatusBadRequest,
			wantBody: &output.DeleteUserResponse{
				Message: "Bad Request",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, body := ui.Delete(tc.arg)
			if status != tc.wantStatus {
				t.Errorf("got %v, want %v", status, tc.wantStatus)
			}
			if body.Message != tc.wantBody.Message {
				t.Errorf("got %v, want %v", body.Message, tc.wantBody.Message)
			}
		})
	}
}
