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
		NameExists(name string) (bool, error)
		NameExistsExcludingId(name string, id uint32) (bool, error)
	}
	MuscleGroupContext interface {
		GetAll() ([]models.MuscleGroup, error)
		GetById(id uint32) (models.MuscleGroup, error)
		GetIdsByNames(names []string) ([]uint32, error)
		Update(id uint32, muscleGroup models.MuscleGroup) error
		Add(muscleGroup models.MuscleGroup) (models.MuscleGroup, error)
		Remove(muscleGroup models.MuscleGroup) error
		NameExists(name string) (bool, error)
		NameExistsExcludingId(name string, id uint32) (bool, error)
		NamesExists(names []string) (bool, error)
	}
	EMGContext interface {
		RemoveAllByExercise(exercise models.Exercise) error
		RemoveAllByMuscleGroup(muscleGroup models.MuscleGroup) error
		SetFromExercise(exerciseId uint32, muscleGroupIds []uint32) error
	}
}
