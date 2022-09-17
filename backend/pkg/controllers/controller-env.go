package controllers

import (
	"github.com/Jaim010/jaim-io/backend/pkg/models"
)

type Env struct {
	ExerciseContext interface {
		GetAll() ([]models.Exercise, error)
		GetById(id int32) (models.Exercise, error)
		Update(id int32, exercis models.Exercise) error
		Add(exercise models.Exercise) (models.Exercise, error)
		Remove(exercise models.Exercise) error
	}
}
