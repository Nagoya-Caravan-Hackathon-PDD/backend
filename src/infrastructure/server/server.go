package server

import (
	"database/sql"

	firebase "firebase.google.com/go"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/infrastructure/http/router"
)

// TODO: サーバの起動と停止を行う
type httpServer struct {
	db  *sql.DB
	app *firebase.App
}

func NewHTTPserver(db *sql.DB, app *firebase.App) *httpServer {
	return &httpServer{db, app}
}

func (s *httpServer) Run() {
	serv := router.NewRouter(s.db, s.app)
	runWithGracefulShutdown(serv)
}
