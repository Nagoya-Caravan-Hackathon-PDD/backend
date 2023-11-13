package middleware

import "net/http"

// TODO: panic Recovery Middlewareを定義する
func (m *middleware) Recovery(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
