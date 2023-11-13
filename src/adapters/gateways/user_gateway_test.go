package gateways

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/pkg/utils"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
	"github.com/lib/pq"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func createOneUser(t *testing.T) types.CreateUser {
	arg := types.CreateUser{
		UserID:   utils.RandomString(10),
		GitHubID: utils.RandomString(10),
	}
	if err := NewUserGateway(dbconn).Create(arg); err != nil {
		t.Fatal(err)
	}
	return arg
}

func TestCreate(t *testing.T) {
	migrateUp(t)
	defer migrateDown(t)
	g := NewUserGateway(dbconn)

	testCases := []struct {
		name        string
		arg         types.CreateUser
		wantErr     error
		wantErrCode string
	}{
		{
			name: "success",
			arg: types.CreateUser{
				UserID:   "test_user_id",
				GitHubID: "test_github_id",
			},
			wantErr: nil,
		},
		{
			name: "already exists",
			arg: types.CreateUser{
				UserID:   "test_user_id",
				GitHubID: "test_github_id",
			},
			wantErr:     &pq.Error{},
			wantErrCode: pgerrcode.UniqueViolation,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := g.Create(tc.arg)
			if reflect.TypeOf(err) != reflect.TypeOf(tc.wantErr) {
				t.Errorf("got %v, want %s", reflect.TypeOf(err), reflect.TypeOf(tc.wantErr))
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

	testRead(t)
	testDelete(t)
}

func testRead(t *testing.T) {

	g := NewUserGateway(dbconn)

	testCases := []struct {
		name        string
		arg         string
		wantErr     error
		wantErrCode string
	}{
		{
			name:    "success",
			arg:     "test_user_id",
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := g.Read(tc.arg)
			if reflect.TypeOf(err) != reflect.TypeOf(tc.wantErr) {
				t.Errorf("got %v, want %s", reflect.TypeOf(err), reflect.TypeOf(tc.wantErr))
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

func testDelete(t *testing.T) {

	g := NewUserGateway(dbconn)

	testCases := []struct {
		name        string
		arg         string
		wantErr     error
		wantErrCode string
	}{
		{
			name:    "success",
			arg:     "test_user_id",
			wantErr: nil,
		},
		{
			name:        "not found",
			arg:         "not_found_user_id",
			wantErr:     &pq.Error{},
			wantErrCode: pgerrcode.NoDataFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := g.Delete(tc.arg)
			if reflect.TypeOf(err) != reflect.TypeOf(tc.wantErr) {
				t.Errorf("got %v, want %s", reflect.TypeOf(err), reflect.TypeOf(tc.wantErr))
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
