package mocks

import (
	"database/sql"
	"fmt"

	"github.com/Jaim010/jaim-io/backend/pkg/models"
)

type MockMuscleGroupContext struct{}

func (c *MockMuscleGroupContext) GetAll() ([]models.MuscleGroup, error) {
	var mgs = []models.MuscleGroup{
		{Id: 1, Name: "Chest", Description: "Lorem ipsum", ImagePath: "/images/chest"},
		{Id: 2, Name: "Tricep", Description: "Lorem ipsum", ImagePath: "/images/tricep"},
	}
	return mgs, nil
}

func (c *MockMuscleGroupContext) GetById(id uint32) (models.MuscleGroup, error) {
	var mgs = []models.MuscleGroup{
		{Id: 1, Name: "Chest", Description: "Lorem ipsum", ImagePath: "/images/chest"},
		{Id: 2, Name: "Tricep", Description: "Lorem ipsum", ImagePath: "/images/tricep"},
	}

	for _, ex := range mgs {
		if ex.Id == id {
			return ex, nil
		}
	}

	return models.MuscleGroup{}, sql.ErrNoRows
}

func (c *MockMuscleGroupContext) Update(id uint32, MuscleGroup models.MuscleGroup) error {
	var mgs = []models.MuscleGroup{
		{Id: 1, Name: "Chest", Description: "Lorem ipsum", ImagePath: "/images/chest"},
		{Id: 2, Name: "Tricep", Description: "Lorem ipsum", ImagePath: "/images/tricep"},
	}

	for i, ex := range mgs {
		if ex.Id == id {
			mgs[i] = MuscleGroup
			return nil
		}
	}

	return fmt.Errorf("muscle group not found")
}

func (c *MockMuscleGroupContext) Add(MuscleGroup models.MuscleGroup) (models.MuscleGroup, error) {
	var mgs = []models.MuscleGroup{
		{Id: 1, Name: "Chest", Description: "Lorem ipsum", ImagePath: "/images/chest"},
		{Id: 2, Name: "Tricep", Description: "Lorem ipsum", ImagePath: "/images/tricep"},
	}

	for _, ex := range mgs {
		if ex.Id == MuscleGroup.Id {
			return models.MuscleGroup{}, fmt.Errorf("muscle group with id %d already exists", MuscleGroup.Id)
		}
	}

	mgs = append(mgs, MuscleGroup)
	return MuscleGroup, nil
}

func (c *MockMuscleGroupContext) Remove(MuscleGroup models.MuscleGroup) error {
	var mgs = []models.MuscleGroup{
		{Id: 1, Name: "Chest", Description: "Lorem ipsum", ImagePath: "/images/chest"},
		{Id: 2, Name: "Tricep", Description: "Lorem ipsum", ImagePath: "/images/tricep"},
	}

	for i, ex := range mgs {
		if ex.Id == MuscleGroup.Id {
			mgs = append(mgs[:i], mgs[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("muscle group not found")
}
