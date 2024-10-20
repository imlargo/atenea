package main

import (
	"encoding/json"
	"extractor-contenido/core"
	"fmt"
	"net/http"
)

func contenidoHandler(w http.ResponseWriter, r *http.Request) {
	// Permitir cualquier origen
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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
