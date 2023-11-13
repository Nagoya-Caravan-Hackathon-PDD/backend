package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

// TODO: ルーティングを定義する

type router struct {
	db   *sql.DB
	echo *echo.Echo
}

func NewRouter(db *sql.DB) *echo.Echo {
	router := &router{
		db:   db,
		echo: echo.New(),
	}

	router.Health()

	return router.echo
}
