package middleware

import (
	"net/http"
)

// TODO: Middlewareのインターフェースを定義する
type Middleware interface {
	Recovery(h http.Handler) http.Handler
}

type middleware struct {
}

func NewMiddleware() Middleware {
	return &middleware{}
}
