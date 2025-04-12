package api

import (

	// This is required to generate swagger docs
	"github.com/gin-gonic/gin"
	"github.com/imlargo/atenea/cmd/docs"
	"github.com/imlargo/atenea/internal/controllers"
	"github.com/imlargo/atenea/internal/middlewares"
	"github.com/imlargo/atenea/internal/repositories"
	"github.com/imlargo/atenea/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Application struct {
	Config Config
}

func (app *Application) Mount() *gin.Engine {
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

	return router
}

func (app *Application) SetupDocs(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Atenea API"
	docs.SwaggerInfo.Description = "Your tool to find university course information quickly and easily."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	urlSwaggerJson := ginSwagger.URL(app.Config.Addr)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, urlSwaggerJson))
}
