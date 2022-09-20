package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jaim010/jaim-io/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

func (env *Env) GetAllExercises(c *gin.Context) {
	exs, err := env.ExerciseContext.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, exs)
	}
}

func (env *Env) GetExerciseById(c *gin.Context) {
	idStr := c.Param("id")

	idU64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint32(idU64)

	ex, err := env.ExerciseContext.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "exercise not found"})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, ex)
}

func (env *Env) PutExercise(c *gin.Context) {
	var newExercise models.Exercise

	idStr := c.Param("id")

	idU64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint32(idU64)

	if err := c.BindJSON(&newExercise); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id != newExercise.Id {
		err := fmt.Sprintf("URI id: '%d' not equal to exercise id: ''%d'", id, newExercise.Id)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := env.ExerciseContext.Update(id, newExercise); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
