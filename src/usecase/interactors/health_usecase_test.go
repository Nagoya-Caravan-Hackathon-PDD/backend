package interactors

import (
	"testing"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/gateways"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/adapters/presenters"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
)

func TestCheckDB(t *testing.T) {
	hu := NewHealthInteractor(gateways.NewHealthGateway(dbconn), presenters.NewHealthPresenter())

	testCases := []struct {
		name       string
		arg        input.HealthRequest
		wantStatus int
		wantOut    output.HealthResponse
	}{
		{
			name: "success no db check",
			arg: input.HealthRequest{
				CheckDB: false,
			},
			wantStatus: 200,
			wantOut: output.HealthResponse{
				Message: "OK",
			},
		},
		{
			name: "success db check",
			arg: input.HealthRequest{
				CheckDB: true,
			},
			wantStatus: 200,
			wantOut: output.HealthResponse{
				Message: "OK",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, out := hu.CheckDB(tc.arg)
			if status != tc.wantStatus {
				t.Errorf("status: got %v, want %v", status, tc.wantStatus)
			}
			if out.Message != tc.wantOut.Message {
				t.Errorf("out: got %v, want %v", out, tc.wantOut)
			}
		})
	}
}
