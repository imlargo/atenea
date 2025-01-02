package middlewares

import (
	"net/http"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var allowedOrigins = map[string]bool{
			"https://atenea-4v0s.onrender.com":                        true,
			"http://localhost:8080":                                   true,
			"http://localhost:5173":                                   true,
			"https://pegaso.imlargo.dev":                              true,
			"https://pegaso-git-develop-imlargos-projects.vercel.app": true,
			"https://sia-extractor-contenidos.onrender.com":           true,
			"https://salidas-campo.vercel.app":                        true,
		}

		origin := r.Header.Get("Origin")

		if !allowedOrigins[origin] {
			http.Error(w, "CORS policy: This origin is not allowed", http.StatusForbidden)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
