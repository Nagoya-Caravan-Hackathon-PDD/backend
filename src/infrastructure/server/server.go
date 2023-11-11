package server

import "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/infrastructure/http/router"

// TODO: サーバの起動と停止を行う
type httpServer struct {
}

func NewHTTPserver() *httpServer {
	return &httpServer{}
}

func (s *httpServer) Run() {
	router := router.NewRouter()
	runWithGracefulShutdown(router.Mux)
}
