package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/imlargo/atenea/config"
	docs "github.com/imlargo/atenea/docs"
	"github.com/imlargo/atenea/internal/controllers"
	"github.com/imlargo/atenea/internal/middlewares"
	"github.com/imlargo/atenea/internal/repositories"
	"github.com/imlargo/atenea/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db int

func init() {
	config.LoadEnvVariables()
	db = 1
	// db = config.ConnectDB()
	// migrations.AutoMigrateAll(db)
}

// @contact.name imlargo
// @contact.url http://www.swagger.io/support
// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	gin.SetMode(os.Getenv("GIN_MODE"))

	docs.SwaggerInfo.Title = "Go API Master Example"
	docs.SwaggerInfo.Description = "This a project is a example of API Rest in Go using Gin, Gorm & PostgreSQL."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	userRepository := repositories.NewCourseRepository(db)
	userService := services.NewCourseService(userRepository)
	userController := controllers.NewCourseController(userService)

	router := gin.Default()

	router.Use(middlewares.NewCorsMiddleware())

	v1 := router.Group("/v1")

	v1.POST("/users", userController.Create)
	v1.PUT("/users", userController.Update)
	v1.GET("/users", userController.GetAll)
	v1.GET("/users/:id", userController.GetById)
	v1.DELETE("/users/:id", userController.Delete)

	urlSwaggerJson := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, urlSwaggerJson))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
