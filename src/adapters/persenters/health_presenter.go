package persenters

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
)

type healthPresenter struct {
	w http.ResponseWriter
}

func NewHealthPresenter(w http.ResponseWriter) *healthPresenter {
	return &healthPresenter{w}
}

func (p *healthPresenter) Success() error {
	p.w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(p.w).Encode(output.HealthResponse{Message: "OK"}); err != nil {
		http.Error(p.w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (p *healthPresenter) Failed(err error) error {
	log.Println(err)
	http.Error(p.w, "Internal Server Error", http.StatusInternalServerError)
	return nil
}
