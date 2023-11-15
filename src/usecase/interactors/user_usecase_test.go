package interactors

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/pkg/utils"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/gateways"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/presenters"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

func createUser(t *testing.T) types.CreateUser {
	arg := types.CreateUser{
		UserID:   utils.RandomString(10),
		GitHubID: utils.RandomString(10),
	}
	gateways.NewUserGateway(dbconn).Create(arg)
	return arg
}

func TestCreate(t *testing.T) {
	mig.Up()
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
			wantBody:   nil,
		},
		{
			name: "bad request github_id empty",
			arg: types.CreateUser{
				UserID:   "",
				GitHubID: "test_github_id",
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   nil,
		},
		{
			name: "conflict",
			arg: types.CreateUser{
				UserID:   "test_user_id",
				GitHubID: "test_github_id",
			},
			wantStatus: http.StatusConflict,
			wantBody:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, body := ui.Create(tc.arg)
			if status != tc.wantStatus {
				t.Errorf("got %v, want %v", status, tc.wantStatus)
			}
			if reflect.TypeOf(body) != reflect.TypeOf(tc.wantBody) {
				t.Errorf("got %v, want %v", body.Message, tc.wantBody.Message)
			}
		})
	}
	testGetUser(t)
	testDelete(t)
}

func testGetUser(t *testing.T) {
	ui := NewUserInteractor(gateways.NewUserGateway(dbconn), presenters.NewUserPresenter())
	testCases := []struct {
		name       string
		arg        input.GetUser
		wantStatus int
		wantBody   *output.ReadUserResponse
	}{
		{
			name: "success",
			arg: input.GetUser{
				UserID: "test_user_id",
			},
			wantStatus: http.StatusOK,
			wantBody: &output.ReadUserResponse{
				UserID:   "test_user_id",
				GitHubID: "test_github_id",
			},
		},
		{
			name: "bad request uid empty",
			arg: input.GetUser{
				UserID: "",
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   nil,
		},
		{
			name: "bad request uid not found",
			arg: input.GetUser{
				UserID: "test_user_id_not_found",
			},
			wantStatus: http.StatusNotFound,
			wantBody:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, body := ui.Read(tc.arg)
			if status != tc.wantStatus {
				t.Errorf("got %v, want %v", status, tc.wantStatus)
			}
			if reflect.TypeOf(body) != reflect.TypeOf(tc.wantBody) {
				t.Errorf("got %v, want %v", body, tc.wantBody)
			}
		})
	}
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
			wantBody:   nil,
		},
		{
			name: "bad request uid not found",
			arg: input.DeleteUsers{
				UserID: "test_user_id",
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, body := ui.Delete(tc.arg)
			if status != tc.wantStatus {
				t.Errorf("got %v, want %v", status, tc.wantStatus)
			}
			if reflect.TypeOf(body) != reflect.TypeOf(tc.wantBody) {
				t.Errorf("got %v, want %v", body.Message, tc.wantBody.Message)
			}
		})
	}
}
