package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type router struct {
	db   *sql.DB
	echo *echo.Echo
}

func NewRouter(db *sql.DB) *echo.Echo {
	router := &router{
		db:  db,git merge develop
		Mux: http.NewServeMux(),
	}

	router.Health()

	return router.echo
}

func buildChain(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	if len(m) == 0 {
		return h
	}
	return m[0](buildChain(h, m[1:cap(m)]...))
}
