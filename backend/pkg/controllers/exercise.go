package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Jaim010/jaim-io/backend/pkg/models"
	"github.com/Jaim010/jaim-io/backend/pkg/models/dto"
	_ "github.com/Jaim010/jaim-io/backend/pkg/utils/httputils"
	"github.com/Jaim010/jaim-io/backend/pkg/utils/utils"
	"github.com/jackc/pgx/v5"

	"github.com/gin-gonic/gin"
)

const (
	ImagePathBase = "../../assets/images/exercises"
)

// GetAllExercises godoc
// @Summary     Get exercises
// @Description get exercises
// @Tags        exercises
// @Accept      json
// @Produce     json
// @Success     200 {array}   dto.ExerciseDTO
// @Failure     400 {object}  httputils.HTTPError
// @Failure     404 {object}  httputils.HTTPError
// @Failure     500 {object} 	httputils.HTTPError
// @Router      /exercise [get]
func (env *Env) GetAllExercises(c *gin.Context) {
	exs, err := env.ExerciseContext.GetAll()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {

		var exsDto []dto.ExerciseDTO
		for _, ex := range exs {
			data, err := ioutil.ReadFile(ex.ImagePath)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			exsDto = append(exsDto, dto.ExerciseDTO{
				Id:           ex.Id,
				Name:         ex.Name,
				Description:  ex.Description,
				MuscleGroups: ex.MuscleGroups,
				ImageData:    data,
				VideoLink:    ex.VideoLink,
			})
		}

		c.IndentedJSON(http.StatusOK, exsDto)
		return
	}
}

// GetExerciseById godoc
// @Summary     Get exercise
// @Description get exercise by ID
// @Tags        exercises
// @Accept      json
// @Produce     json
// @Param       id  path       int 								 true "Exercise ID" Format(uint32)
// @Success     200 {object} 	 dto.ExerciseDTO
// @Failure     400 {object} 	 httputils.HTTPError
// @Failure     404 {object} 	 httputils.HTTPError
// @Failure     500 {object} 	 httputils.HTTPError
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
		if err == pgx.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "exercise not found"})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	data, err := ioutil.ReadFile(ex.ImagePath)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	exDto := dto.ExerciseDTO{
		Id:           ex.Id,
		Name:         ex.Name,
		Description:  ex.Description,
		MuscleGroups: ex.MuscleGroups,
		ImageData:    data,
		VideoLink:    ex.VideoLink,
	}

	c.IndentedJSON(http.StatusOK, exDto)
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
// @Failure     400 			{object} httputils.HTTPError
// @Failure     500 			{object} httputils.HTTPError
// @Router      /exercise/{id} [put]
func (env *Env) PutExercise(c *gin.Context) {
	var exerciseDto dto.ExerciseDTO

	idStr := c.Param("id")

	id, err := utils.StrToUint32(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&exerciseDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id != exerciseDto.Id {
		err := fmt.Sprintf("URI id: '%d' not equal to exercise id: ''%d'", id, exerciseDto.Id)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Checks if exercise exists
	// If exercise already exists, reject request
	ex_exists, err := env.ExerciseContext.NameExistsExcludingId(exerciseDto.Name, exerciseDto.Id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ex_exists {
		err := fmt.Sprintf("an exercise with name '%s' already exists", exerciseDto.Name)
		c.IndentedJSON(http.StatusConflict, gin.H{"error": err})
		return
	}

	// Checks if all muscle groups exist
	// If not rejects request
	mgs_exist, err := env.MuscleGroupContext.NamesExists(exerciseDto.MuscleGroups)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !mgs_exist {
		err := fmt.Sprintf("one or more muscle groups do not exist: %v", exerciseDto.MuscleGroups)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	imagePath := fmt.Sprintf("%s/%d", ImagePathBase, exerciseDto.Id)
	updatedExercise := models.Exercise{
		Id:           exerciseDto.Id,
		Name:         exerciseDto.Name,
		Description:  exerciseDto.Description,
		MuscleGroups: exerciseDto.MuscleGroups,
		ImagePath:    imagePath,
		VideoLink:    exerciseDto.VideoLink,
	}

	oldImage, err := ioutil.ReadFile(imagePath)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Checks if the file has changed
	if !bytes.Equal(oldImage, exerciseDto.ImageData) {
		os.Remove(imagePath)
		if len(exerciseDto.ImageData) != 0 {
			os.WriteFile(updatedExercise.ImagePath, exerciseDto.ImageData, 0664)
		}
	}

	// Removes all the current exercise_muscle_groups relations, related to the exercise
	if err := env.EMGContext.RemoveAllByExercise(updatedExercise); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "RemoveUnusedByExercise: " + err.Error()})
		return
	}

	// Gets all the muscle_groups IDs of the updated exercise
	ids, err := env.MuscleGroupContext.GetIdsByNames(updatedExercise.MuscleGroups)
	if err != nil && err != pgx.ErrNoRows {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "GetIdsByNames: " + err.Error()})
		return
	}

	// Sets all new relations
	if err := env.EMGContext.SetFromExercise(updatedExercise.Id, ids); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "SetFromExercise: " + err.Error()})
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
// @Failure     400 			{object} httputils.HTTPError
// @Failure     500 			{object} httputils.HTTPError
// @Router      /exercise/ [post]
func (env *Env) PostExercise(c *gin.Context) {
	var exerciseDto dto.ExerciseDTO

	if err := c.BindJSON(&exerciseDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if empty := exerciseDto.Description == "" || exerciseDto.VideoLink == "" || len(exerciseDto.MuscleGroups) == 0; empty {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "exercise fields can't be completely empty"})
		return
	}

	// Checks if exercise exists
	// If exercise already exists, reject request
	ex_exists, err := env.ExerciseContext.NameExists(exerciseDto.Name)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ex_exists {
		err := fmt.Sprintf("an exercise with name '%s' already exists", exerciseDto.Name)
		c.IndentedJSON(http.StatusConflict, gin.H{"error": err})
		return
	}

	// Checks if all muscle groups exist
	// If not rejects request
	mgs_exist, err := env.MuscleGroupContext.NamesExists(exerciseDto.MuscleGroups)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !mgs_exist {
		err := fmt.Sprintf("one or more muscle groups do not exist: %v", exerciseDto.MuscleGroups)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	imagePath := fmt.Sprintf("%s/%d", ImagePathBase, exerciseDto.Id)
	newExercise := models.Exercise{
		Id:           exerciseDto.Id,
		Name:         exerciseDto.Name,
		Description:  exerciseDto.Description,
		MuscleGroups: exerciseDto.MuscleGroups,
		ImagePath:    imagePath,
		VideoLink:    exerciseDto.VideoLink,
	}
	if len(exerciseDto.ImageData) != 0 {
		os.WriteFile(newExercise.ImagePath, exerciseDto.ImageData, 0664)
	}

	exercise, err := env.ExerciseContext.Add(newExercise)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Gets all the muscle_groups IDs of the updated exercise
	ids, err := env.MuscleGroupContext.GetIdsByNames(exerciseDto.MuscleGroups)
	if err != nil && err != pgx.ErrNoRows {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "GetIdsByNames: " + err.Error()})
		return
	}

	// Sets all new relations
	if err := env.EMGContext.SetFromExercise(exerciseDto.Id, ids); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "SetFromExercise: " + err.Error()})
		return
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
// @Failure     400 			{object} httputils.HTTPError
// @Failure     404 			{object} httputils.HTTPError
// @Failure     500 			{object} httputils.HTTPError
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
		if err == pgx.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "exercise not found"})
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = env.EMGContext.RemoveAllByExercise(exercise)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = env.ExerciseContext.Remove(exercise)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	imagePath := fmt.Sprintf("%s/%d", ImagePathBase, id)
	os.Remove(imagePath)

	c.Status(http.StatusNoContent)
}
