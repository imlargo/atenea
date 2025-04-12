package middlewares

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCorsMiddleware() gin.HandlerFunc {

	allowedOrigins := []string{
		"https://atenea-4v0s.onrender.com",
		"http://localhost:8080",
		"http://localhost:5173",
		"https://pegaso.imlargo.dev",
		"https://pegaso-git-develop-imlargos-projects.vercel.app",
		"https://sia-extractor-contenidos.onrender.com",
		"https://salidas-campo.vercel.app",
		"https://repo-contenidos-minas.vercel.app",
		"http://localhost:4173",
		"https://atenea-un.vercel.app",
	}

	config := cors.Config{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodHead,
			http.MethodOptions,
		},
		AllowCredentials: true,
	}

	return cors.New(config)
}
