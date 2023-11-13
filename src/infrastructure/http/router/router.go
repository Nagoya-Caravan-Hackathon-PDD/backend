package router

import (
	"database/sql"
	"net/http"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/infrastructure/http/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type router struct {
	db         *sql.DB
	echo       *echo.Echo
	middleware middleware.Middleware
}

func NewRouter(db *sql.DB) *echo.Echo {
	router := &router{
		db:         db,
		echo:       echo.New(),
		middleware: middleware.NewMiddleware(),
	}

	router.echo.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{},
	}), echoMiddleware.Recover())

	router.Health()
	router.userRouter()
	router.encounterRoutes()

	return router.echo
}
