package controllers

import (
	"fmt"
	"net/http"

	"github.com/Jaim010/jaim-io/backend/pkg/models"
	_ "github.com/Jaim010/jaim-io/backend/pkg/utils/httputils"
	"github.com/Jaim010/jaim-io/backend/pkg/utils/utils"
	"github.com/jackc/pgx/v5"

	"github.com/gin-gonic/gin"
)

// GetAllMuscleGroups godoc
// @Summary     Get muscle groups
// @Description get muscle groups
// @Tags        muscle_groups
// @Accept      json
// @Produce     json
// @Success     200 {array}   	models.MuscleGroup
// @Failure     400 {object}   	httputil.HTTPError
// @Failure     404 {object}  	httputil.HTTPError
// @Failure     500 {object} 		httputil.HTTPError
// @Router      /musclegroup [get]
func (env *Env) GetAllMuscleGroups(c *gin.Context) {
	mgs, err := env.MuscleGroupContext.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.IndentedJSON(http.StatusOK, mgs)
		return
	}
}

// GetMuscleGroupById godoc
// @Summary     Get muscle group
// @Description get muscle group by ID
// @Tags        muscle_groups
// @Accept      json
// @Produce     json
// @Param       id  path       int 								 true "Muscle group ID" Format(uint32)
// @Success     200 {object} 	 models.MuscleGroup
// @Failure     400 {object} 	 httputil.HTTPError
// @Failure     404 {object} 	 httputil.HTTPError
// @Failure     500 {object} 	 httputil.HTTPError
// @Router      /musclegroup/{id} [get]
func (env *Env) GetMuscleGroupById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := utils.StrToUint32(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mg, err := env.MuscleGroupContext.GetById(id)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "exercise not found"})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.IndentedJSON(http.StatusOK, mg)
}

// PutMuscleGroup godoc
// @Summary     Update muscle group
// @Description update by json muscle group
// @Tags        muscle_groups
// @Accept      json
// @Produce     json
// @Param       id  			path     int 								 true "Muscle group ID" Format(uint32)
// @Param       exercise  body     models.MuscleGroup		 true "Update muscle group"
// @Success     204
// @Failure     400 			{object} httputil.HTTPError
// @Failure     500 			{object} httputil.HTTPError
// @Router      /musclegroup/{id} [put]
func (env *Env) PutMuscleGroup(c *gin.Context) {
	var updatedMuscleGroup models.MuscleGroup

	idStr := c.Param("id")
	id, err := utils.StrToUint32(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&updatedMuscleGroup); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id != updatedMuscleGroup.Id {
		err := fmt.Sprintf("URI id: '%d' not equal to muscle group id: ''%d'", id, updatedMuscleGroup.Id)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Checks if muscle group exists
	// If muscle group already exists, reject request
	mg_exists, err := env.MuscleGroupContext.NameExistsExcludingId(updatedMuscleGroup.Name, updatedMuscleGroup.Id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if mg_exists {
		err := fmt.Sprintf("another muscle group with name '%s' already exists", updatedMuscleGroup.Name)
		c.IndentedJSON(http.StatusConflict, gin.H{"error": err})
		return
	}

	if err := env.MuscleGroupContext.Update(id, updatedMuscleGroup); err != nil {
		if _, err := env.MuscleGroupContext.GetById(id); err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// PostMuscleGroup godoc
// @Summary     Add muscle group
// @Description add by json muscle group
// @Tags        muscle_groups
// @Accept      json
// @Produce     json
// @Param       musclegroup  body     models.MuscleGroup		 true "Add muscle group"
// @Success     201				{object} models.MuscleGroup
// @Failure     400 			{object} httputil.HTTPError
// @Failure     500 			{object} httputil.HTTPError
// @Router      /musclegroup/ [post]
func (env *Env) PostMuscleGroup(c *gin.Context) {
	var newMuscleGroup models.MuscleGroup

	if err := c.BindJSON(&newMuscleGroup); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if empty := newMuscleGroup.Description == "" || newMuscleGroup.ImagePath == ""; empty {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "muscle group fields can't be completely empty"})
		return
	}

	// Checks if muscle group exists
	// If muscle group already exists, reject request
	mg_exists, err := env.MuscleGroupContext.NameExists(newMuscleGroup.Name)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if mg_exists {
		err := fmt.Sprintf("a muscle group with name '%s' already exists", newMuscleGroup.Name)
		c.IndentedJSON(http.StatusConflict, gin.H{"error": err})
		return
	}

	exercise, err := env.MuscleGroupContext.Add(newMuscleGroup)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, exercise)
}

// DeleteMuscleGroup godoc
// @Summary     Delete an muscle group
// @Description delete by muscle group id
// @Tags        muscle_groups
// @Accept      json
// @Produce     json
// @Param       id  			path     int 								 true "Muscle group ID" Format(uint32)
// @Success     204
// @Failure     400 			{object} httputil.HTTPError
// @Failure     404 			{object} httputil.HTTPError
// @Failure     500 			{object} httputil.HTTPError
// @Router      /musclegroup/{id} [delete]
func (env *Env) DeleteMuscleGroup(c *gin.Context) {
	idStr := c.Param("id")

	id, err := utils.StrToUint32(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	muscleGroup, err := env.MuscleGroupContext.GetById(id)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "muscle group not found"})
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = env.EMGContext.RemoveAllByMuscleGroup(muscleGroup)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = env.MuscleGroupContext.Remove(muscleGroup)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
