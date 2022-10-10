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
		Exists(name string) (bool, error)
		ExistsExcludingId(name string, id uint32) (bool, error)
	}
	MuscleGroupContext interface {
		GetAll() ([]models.MuscleGroup, error)
		GetById(id uint32) (models.MuscleGroup, error)
		Update(id uint32, muscleGroup models.MuscleGroup) error
		Add(muscleGroup models.MuscleGroup) (models.MuscleGroup, error)
		Remove(muscleGroup models.MuscleGroup) error
		Exists(name string) (bool, error)
		ExistsExcludingId(name string, id uint32) (bool, error)
		ExistsArr(names []string) (bool, error)
	}
}
