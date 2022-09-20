package main

import (
	"log"

	"github.com/Jaim010/jaim-io/backend/pkg/controllers"
	database "github.com/Jaim010/jaim-io/backend/pkg/db"
	"github.com/Jaim010/jaim-io/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

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
	api := router.Group("/api")
	{
		api.GET("/exercises", env.GetAllExercises)
	}
	router.Run("0.0.0.0:8080")
}
