package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/Jaim010/jaim-io/backend/pkg/httputil"
	"github.com/Jaim010/jaim-io/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// GetAllExercises godoc
// @Summary     Get exercises
// @Description get exercises
// @Tags        exercises
// @Accept      json
// @Produce     json
// @Success     200 {array}   models.Exercise
// @Failure     400 {array}   httputil.HTTPError
// @Failure     404 {array}  httputil.HTTPError
// @Failure     500 {object} httputil.HTTPError
// @Router      /exercise [get]
func (env *Env) GetAllExercises(c *gin.Context) {
	exs, err := env.ExerciseContext.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, exs)
	}
}

// GetExerciseById godoc
// @Summary     Get exercise
// @Description get exervise by ID
// @Tags        exercises
// @Accept      json
// @Produce     json
// @Param       id  path       int true "Exercise ID"
// @Success     200 {object} models.Exercise
// @Failure     400 {object} httputil.HTTPError
// @Failure     404 {object} httputil.HTTPError
// @Failure     500 {object} httputil.HTTPError
// @Router      /exercise/{id} [get]
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
