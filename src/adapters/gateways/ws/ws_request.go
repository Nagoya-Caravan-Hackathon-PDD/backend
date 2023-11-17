package ws

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

type WSRequest struct {
}

func NewWSRequest() *WSRequest {
	return &WSRequest{}
}

func (r *WSRequest) Start(times int, token string) error {
	return r.timer(times, token, types.FlagStart)
}

func (r *WSRequest) Turn(times int, token string) error {
	return r.timer(times, token, types.FlagTurn)
}

func (r *WSRequest) Result(times int, token string) error {
	return r.timer(times, token, types.FlagResult)
}

func (r *WSRequest) End(times int, token string) error {
	return r.timer(times, token, types.FlagEnd)
}

func (r *WSRequest) timer(times int, token string, flag types.MessageType) error {
	request := types.WSRequest{
		MessageType: flag,
		Time:        times,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, config.Config.Server.WSURL+token, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println(resp.StatusCode)

	return nil
}
