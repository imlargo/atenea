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

func init() {
	config.LoadEnvVariables()
}

// @contact.name imlargo
// @contact.url http://www.swagger.io/support
// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	gin.SetMode(os.Getenv("GIN_MODE"))

	docs.SwaggerInfo.Title = "Atenea API"
	docs.SwaggerInfo.Description = "Your tool to find university course information quickly and easily."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	courseRepository := repositories.NewCourseRepository(1)
	courseService := services.NewCourseService(courseRepository)
	courseController := controllers.NewCourseController(courseService)

	router := gin.Default()

	router.Use(middlewares.NewCorsMiddleware())

	v1 := router.Group("/v1")

	v1.POST("/courses", courseController.Create)
	v1.PUT("/courses", courseController.Update)
	v1.GET("/courses", courseController.GetAll)
	v1.GET("/courses/:id", courseController.GetById)
	v1.DELETE("/courses/:id", courseController.Delete)

	urlSwaggerJson := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, urlSwaggerJson))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
