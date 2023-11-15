package middleware

import "github.com/labstack/echo/v4"

// TODO: Middlewareのインターフェースを定義する
type Middleware interface {
	FirebaseAuth(n echo.HandlerFunc) echo.HandlerFunc
}

type middleware struct {
}

func NewMiddleware() Middleware {
	return &middleware{}
}
