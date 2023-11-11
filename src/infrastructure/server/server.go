package server

import (
	"database/sql"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/infrastructure/http/router"
)

// TODO: サーバの起動と停止を行う
type httpServer struct {
	db *sql.DB
}

func NewHTTPserver(db *sql.DB) *httpServer {
	return &httpServer{db}
}

func (s *httpServer) Run() {
	router := router.NewRouter(s.db)
	runWithGracefulShutdown(router.Mux)
}
