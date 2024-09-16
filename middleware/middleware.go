package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json") //adicionando um content type para o navegador identificar como JSON
		next.ServeHTTP(w, r)
	})
}
