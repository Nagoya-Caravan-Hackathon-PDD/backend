package router

import "net/http"

// TODO: ルーティングを定義する

type router struct {
	Mux *http.ServeMux
}

func NewRouter() *router {
	router := &router{
		Mux: http.NewServeMux(),
	}

	return router
}
