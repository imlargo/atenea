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
		Port:   os.Getenv(env.PORT),
		ApiURL: os.Getenv(env.API_URL),
	}

	app := &api.Application{
		Config: config,
	}

	router := app.Mount()

	app.SetupDocs(router)

	if err := router.Run(":" + app.Config.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
