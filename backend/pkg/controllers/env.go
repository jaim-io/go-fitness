package controllers

import (
	"github.com/Jaim010/jaim-io/backend/pkg/models"
)

type Env struct {
	ExerciseContext interface {
		GetAll() ([]models.Exercise, error)
		GetById(id uint32) (models.Exercise, error)
		Update(id uint32, exercis models.Exercise) error
		Add(exercise models.Exercise) (models.Exercise, error)
		Remove(exercise models.Exercise) error
	}
}
