package controllers

import (
	"log"
	"net/http"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/usecase/ports"
	"github.com/bytedance/go-tagexpr/v2/binding"
)

type healthController struct {
	usecase ports.HealthInput
}

func NewHealthController(usecase ports.HealthInput) *healthController {
	return &healthController{usecase}
}

func (h *healthController) Health(w http.ResponseWriter, r *http.Request) {
	var reqQuery input.HealthRequest

	if err := binding.New(nil).Bind(&reqQuery, r, nil); err != nil {
		log.Println("failed to bind and validate request :", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	h.usecase.CheckDB(reqQuery)
}
