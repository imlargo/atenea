package routes

import (
	"atenea/src/services/extractor"
	"encoding/json"
	"net/http"
)

func GetContentByCode(w http.ResponseWriter, r *http.Request) {
	// Manejar solicitudes OPTIONS
	if r.Method == http.MethodOptions {
		return
	}

	codigo := r.URL.Query().Get("codigo")
	if codigo == "" {
		http.Error(w, "codigo parameter is missing", http.StatusBadRequest)
		return
	}

	asignatura := extractor.GetContenidoAsignatura(codigo)
	if asignatura == nil {
		http.Error(w, "No content found for the given codigo", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(asignatura)
}
