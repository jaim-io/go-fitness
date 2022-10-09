package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/Jaim010/jaim-io/backend/pkg/httputil"
	"github.com/Jaim010/jaim-io/backend/pkg/models"
	"github.com/Jaim010/jaim-io/backend/pkg/utils/utils"

	"github.com/gin-gonic/gin"
)

// GetAllExercises godoc
// @Summary     Get exercises
// @Description get exercises
// @Tags        exercises
// @Accept      json
// @Produce     json
// @Success     200 {array}   models.Exercise
// @Failure     400 {object}   httputil.HTTPError
// @Failure     404 {object}  	httputil.HTTPError
// @Failure     500 {object} 	httputil.HTTPError
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
// @Description get exercise by ID
// @Tags        exercises
// @Accept      json
// @Produce     json
// @Param       id  path       int 								 true "Exercise ID" Format(uint32)
// @Success     200 {object} 	 models.Exercise
// @Failure     400 {object} 	 httputil.HTTPError
// @Failure     404 {object} 	 httputil.HTTPError
// @Failure     500 {object} 	 httputil.HTTPError
// @Router      /exercise/{id} [get]
func (env *Env) GetExerciseById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := utils.StrToUint32(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

// PutExercise godoc
// @Summary     Update exercise
// @Description update by json exercise
// @Tags        exercises
// @Accept      json
// @Produce     json
// @Param       id  			path     int 								 true "Exercise ID" Format(uint32)
// @Param       exercise  body     models.Exercise		 true "Update exercise"
// @Success     204
// @Failure     400 			{object} httputil.HTTPError
// @Failure     500 			{object} httputil.HTTPError
// @Router      /exercise/{id} [put]
func (env *Env) PutExercise(c *gin.Context) {
	var updatedExercise models.Exercise

	idStr := c.Param("id")

	id, err := utils.StrToUint32(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&updatedExercise); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id != updatedExercise.Id {
		err := fmt.Sprintf("URI id: '%d' not equal to exercise id: ''%d'", id, updatedExercise.Id)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := env.ExerciseContext.Update(id, updatedExercise); err != nil {
		if _, err := env.ExerciseContext.GetById(id); err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// PostExercise godoc
// @Summary     Add exercise
// @Description add by json exercise
// @Tags        exercises
// @Accept      json
// @Produce     json
// @Param       exercise  body     models.Exercise		 true "Add exercise"
// @Success     201				{object} models.Exercise
// @Failure     400 			{object} httputil.HTTPError
// @Failure     500 			{object} httputil.HTTPError
// @Router      /exercise/ [post]
func (env *Env) PostExercise(c *gin.Context) {
	var newExercise models.Exercise

	if err := c.BindJSON(&newExercise); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if empty := newExercise.Description == "" || newExercise.ImagePath == "" || newExercise.VideoLink == "" || len(newExercise.MuscleGroups) == 0; empty {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "exercise fields can't be completely empty"})
		return
	}

	exercise, err := env.ExerciseContext.Add(newExercise)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusCreated, exercise)
}

// DeleteExercise godoc
// @Summary     Delete an exercise
// @Description delete by exercise id
// @Tags        exercises
// @Accept      json
// @Produce     json
// @Param       id  			path     int 								 true "Exercise ID" Format(uint32)
// @Success     204
// @Failure     400 			{object} httputil.HTTPError
// @Failure     404 			{object} httputil.HTTPError
// @Failure     500 			{object} httputil.HTTPError
// @Router      /exercise/{id} [delete]
func (env *Env) DeleteExercise(c *gin.Context) {
	idStr := c.Param("id")

	id, err := utils.StrToUint32(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercise, err := env.ExerciseContext.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "exercise not found"})
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = env.ExerciseContext.Remove(exercise)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
