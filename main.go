package main

import (
	"encoding/json"
	"extractor-contenido/core"
	"fmt"
	"net/http"
)

func contenidoHandler(w http.ResponseWriter, r *http.Request) {
	// Lista de dominios permitidos
	allowedOrigins := []string{
		"http://localhost:5173",
		"https://pegaso.imlargo.dev",
		"https://pegaso-git-develop-imlargos-projects.vercel.app",
		"https://sia-extractor-contenidos.onrender.com",
	}

	origin := r.Header.Get("Origin")
	allowed := false
	for _, o := range allowedOrigins {
		if o == origin {
			allowed = true
			break
		}
	}

	if allowed {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	} else {
		http.Error(w, "CORS policy: This origin is not allowed", http.StatusForbidden)
		return
	}

	// Manejar solicitudes OPTIONS
	if r.Method == http.MethodOptions {
		return
	}

	codigo := r.URL.Query().Get("codigo")
	if codigo == "" {
		http.Error(w, "codigo parameter is missing", http.StatusBadRequest)
		return
	}

	asignatura := core.GetContenidoAsignatura(codigo)
	if asignatura == nil {
		http.Error(w, "No content found for the given codigo", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(asignatura)
}

func main() {
	http.HandleFunc("/contenido", contenidoHandler)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
