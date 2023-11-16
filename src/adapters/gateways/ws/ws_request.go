package ws

import (
	"bytes"
	"encoding/json"
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
	request := types.WSRequest{
		MessageType: types.FlagStart,
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

	if resp.StatusCode != http.StatusOK {
		return types.ErrBadResponse
	}

	return nil
}
