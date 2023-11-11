package router

import (
	"database/sql"
	"net/http"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/infrastructure/http/middleware"
)

type router struct {
	db         *sql.DB
	Mux        *http.ServeMux
	middleware middleware.Middleware
}

func NewRouter(db *sql.DB) *router {
	router := &router{
		db:         db,
		Mux:        http.NewServeMux(),
		middleware: middleware.NewMiddleware(),
	}

	router.Health()

	return router
}

func buildChain(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	if len(m) == 0 {
		return h
	}
	return m[0](buildChain(h, m[1:cap(m)]...))
}
