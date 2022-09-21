package main

import (
	"log"

	_ "github.com/Jaim010/jaim-io/backend/docs"
	"github.com/Jaim010/jaim-io/backend/pkg/controllers"
	database "github.com/Jaim010/jaim-io/backend/pkg/db"
	"github.com/Jaim010/jaim-io/backend/pkg/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title       Jaim-io
// @version     1.0
// @description My portfolio website

// @contact.name  Jamey Schaap
// @contact.url   https://www.linkedin.com/in/jamey-schaap/
// @contact.email jameyschaap06@gmail.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /api/v1
func main() {

	db, err := database.Init()
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %s", err)
	}

	env := &controllers.Env{
		ExerciseContext: models.ExerciseContext{DB: db},
	}

	router := gin.Default()

	router.GET("/health", controllers.GetHealth)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api/v1")
	{
		api.GET("/exercise", env.GetAllExercises)
		api.GET("/exercise/:id", env.GetExerciseById)
	}

	router.Run("0.0.0.0:8080")
}
