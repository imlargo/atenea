package app

import (
	"atenea/src/middlewares"
	"atenea/src/routes"
	"fmt"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/contenido", routes.GetContentByCode)
}

func SetupServer() http.Handler {
	mux := http.NewServeMux()

	SetupRoutes(mux)

	// Envuelve el manejador principal con el middleware Cors
	handler := middlewares.Cors(mux)

	return handler
}

func StartServer() {

	handler := SetupServer()

	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
