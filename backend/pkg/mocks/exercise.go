package mocks

import (
	"database/sql"
	"fmt"

	"github.com/Jaim010/jaim-io/backend/pkg/models"
)

type MockExerciseContext struct{}

func (c *MockExerciseContext) GetAll() ([]models.Exercise, error) {
	var exs = []models.Exercise{
		{Name: "Barbell bench press", Id: 1, Description: "Lorem ipsum"},
		{Name: "Bulgarian split squat", Id: 2, Description: "Lorem ipsum"},
	}
	return exs, nil
}

func (c *MockExerciseContext) GetById(id uint32) (models.Exercise, error) {
	var exs = []models.Exercise{
		{Name: "Barbell bench press", Id: 1, Description: "Lorem ipsum"},
		{Name: "Bulgarian split squat", Id: 2, Description: "Lorem ipsum"},
	}

	for _, ex := range exs {
		if ex.Id == id {
			return ex, nil
		}
	}

	return models.Exercise{}, sql.ErrNoRows
}

func (c *MockExerciseContext) Update(id uint32, exercise models.Exercise) error {
	var exs = []models.Exercise{
		{Name: "Barbell bench press", Id: 1, Description: "Lorem ipsum"},
		{Name: "Bulgarian split squat", Id: 2, Description: "Lorem ipsum"},
	}

	for i, ex := range exs {
		if ex.Id == id {
			exs[i] = exercise
			return nil
		}
	}

	return fmt.Errorf("Exercise not found")
}

func (c *MockExerciseContext) Add(exercise models.Exercise) (models.Exercise, error) {
	var exs = []models.Exercise{
		{Name: "Barbell bench press", Id: 1, Description: "Lorem ipsum"},
		{Name: "Bulgarian split squat", Id: 2, Description: "Lorem ipsum"},
	}

	for _, ex := range exs {
		if ex.Id == exercise.Id {
			return models.Exercise{}, fmt.Errorf("Exercise with id %d already exists", exercise.Id)
		}
	}

	exs = append(exs, exercise)
	return exercise, nil
}

func (c *MockExerciseContext) Remove(exercise models.Exercise) error {
	var exs = []models.Exercise{
		{Name: "Barbell bench press", Id: 1, Description: "Lorem ipsum"},
		{Name: "Bulgarian split squat", Id: 2, Description: "Lorem ipsum"},
	}

	for i, ex := range exs {
		if ex.Id == exercise.Id {
			exs = append(exs[:i], exs[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Exercise not found")
}
