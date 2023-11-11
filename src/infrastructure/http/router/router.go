package router

import (
	"database/sql"
	"net/http"
)

// TODO: ルーティングを定義する

type router struct {
	db  *sql.DB
	Mux *http.ServeMux
}

func NewRouter(db *sql.DB) *router {
	router := &router{
		db:  db,
		Mux: http.NewServeMux(),
	}

	router.Health()

	return router
}
