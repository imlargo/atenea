package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/imlargo/atenea/cmd/api"
	"github.com/imlargo/atenea/internal/env"
)

// @contact.name imlargo
// @contact.url http://www.swagger.io/support
// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	env.Initialize()

	gin.SetMode(os.Getenv("GIN_MODE"))

	config := api.Config{
		Addr:   ":8080",
		ApiURL: "http://localhost:8080",
	}

	app := &api.Application{
		Config: config,
	}

	router := app.Mount()

	app.SetupDocs(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
