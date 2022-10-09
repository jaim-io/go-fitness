package mocks

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Jaim010/jaim-io/backend/pkg/models"
)

type MockExerciseContext struct{}

func (c *MockExerciseContext) GetAll() ([]models.Exercise, error) {
	var exs = []models.Exercise{
		{Id: 1, Name: "Barbell bench press", MuscleGroups: []string{"Chest", "Tricep"}, Description: "Lorem ipsum", ImagePath: "/images/bb_bench_press", VideoLink: "https://www.youtube.com/"},
		{Id: 2, Name: "Bulgarian split squat", MuscleGroups: []string{"Quad", "Glute"}, Description: "Lorem ipsum", ImagePath: "/images/b_split_squad", VideoLink: "https://www.youtube.com/"},
	}
	return exs, nil
}

func (c *MockExerciseContext) GetById(id uint32) (models.Exercise, error) {
	var exs = []models.Exercise{
		{Id: 1, Name: "Barbell bench press", MuscleGroups: []string{"Chest", "Tricep"}, Description: "Lorem ipsum", ImagePath: "/images/bb_bench_press", VideoLink: "https://www.youtube.com/"},
		{Id: 2, Name: "Bulgarian split squat", MuscleGroups: []string{"Quad", "Glute"}, Description: "Lorem ipsum", ImagePath: "/images/b_split_squad", VideoLink: "https://www.youtube.com/"},
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
		{Id: 1, Name: "Barbell bench press", MuscleGroups: []string{"Chest", "Tricep"}, Description: "Lorem ipsum", ImagePath: "/images/bb_bench_press", VideoLink: "https://www.youtube.com/"},
		{Id: 2, Name: "Bulgarian split squat", MuscleGroups: []string{"Quad", "Glute"}, Description: "Lorem ipsum", ImagePath: "/images/b_split_squad", VideoLink: "https://www.youtube.com/"},
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
		{Id: 1, Name: "Barbell bench press", MuscleGroups: []string{"Chest", "Tricep"}, Description: "Lorem ipsum", ImagePath: "/images/bb_bench_press", VideoLink: "https://www.youtube.com/"},
		{Id: 2, Name: "Bulgarian split squat", MuscleGroups: []string{"Quad", "Glute"}, Description: "Lorem ipsum", ImagePath: "/images/b_split_squad", VideoLink: "https://www.youtube.com/"},
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
		{Id: 1, Name: "Barbell bench press", MuscleGroups: []string{"Chest", "Tricep"}, Description: "Lorem ipsum", ImagePath: "/images/bb_bench_press", VideoLink: "https://www.youtube.com/"},
		{Id: 2, Name: "Bulgarian split squat", MuscleGroups: []string{"Quad", "Glute"}, Description: "Lorem ipsum", ImagePath: "/images/b_split_squad", VideoLink: "https://www.youtube.com/"},
	}

	for i, ex := range exs {
		if ex.Id == exercise.Id {
			exs = append(exs[:i], exs[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Exercise not found")
}

func (c *MockExerciseContext) Exists(name string) (bool, error) {
	var exs = []models.Exercise{
		{Id: 1, Name: "Barbell bench press", MuscleGroups: []string{"Chest", "Tricep"}, Description: "Lorem ipsum", ImagePath: "/images/bb_bench_press", VideoLink: "https://www.youtube.com/"},
		{Id: 2, Name: "Bulgarian split squat", MuscleGroups: []string{"Quad", "Glute"}, Description: "Lorem ipsum", ImagePath: "/images/b_split_squad", VideoLink: "https://www.youtube.com/"},
	}

	low_name := strings.ToLower(name)
	for _, ex := range exs {
		if strings.ToLower(ex.Name) == low_name {
			return true, nil
		}
	}
	return false, nil
}
