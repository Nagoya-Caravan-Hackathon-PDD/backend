package middleware

// TODO: Middlewareのインターフェースを定義する
type Middleware interface {
}

type middleware struct {
}

func NewMiddleware() Middleware {
	return &middleware{}
}
