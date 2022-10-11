package mocks

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Jaim010/jaim-io/backend/pkg/models"
	"github.com/jackc/pgx/v5"
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

func (c *MockMuscleGroupContext) NameExists(name string) (bool, error) {
	var mgs = []models.MuscleGroup{
		{Id: 1, Name: "Chest", Description: "Lorem ipsum", ImagePath: "/images/chest"},
		{Id: 2, Name: "Tricep", Description: "Lorem ipsum", ImagePath: "/images/tricep"},
	}

	low_name := strings.ToLower(name)
	for _, mg := range mgs {
		if strings.ToLower(mg.Name) == low_name {
			return true, nil
		}
	}
	return false, nil
}

func (c *MockMuscleGroupContext) NameExistsExcludingId(name string, id uint32) (bool, error) {
	var mgs = []models.MuscleGroup{
		{Id: 1, Name: "Chest", Description: "Lorem ipsum", ImagePath: "/images/chest"},
		{Id: 2, Name: "Tricep", Description: "Lorem ipsum", ImagePath: "/images/tricep"},
	}

	low_name := strings.ToLower(name)
	for _, mg := range mgs {
		if strings.ToLower(mg.Name) == low_name && mg.Id != uint32(id) {
			return true, nil
		}
	}
	return false, nil
}

func (c *MockMuscleGroupContext) NamesExists(names []string) (bool, error) {
	var mgs = []models.MuscleGroup{
		{Id: 1, Name: "Chest", Description: "Lorem ipsum", ImagePath: "/images/chest"},
		{Id: 2, Name: "Tricep", Description: "Lorem ipsum", ImagePath: "/images/tricep"},
	}

	for _, name := range names {
		low_name := strings.ToLower(name)
		for _, mg := range mgs {
			if strings.ToLower(mg.Name) == low_name {
				return true, nil
			}
		}
	}
	return false, nil
}

func (c *MockMuscleGroupContext) GetIdsByNames(names []string) ([]uint32, error) {
	var mgs = []models.MuscleGroup{
		{Id: 1, Name: "Chest", Description: "Lorem ipsum", ImagePath: "/images/chest"},
		{Id: 2, Name: "Tricep", Description: "Lorem ipsum", ImagePath: "/images/tricep"},
	}

	var ids []uint32

	for _, name := range names {
		low_name := strings.ToLower(name)
		for _, mg := range mgs {
			if mg.Name == low_name {
				ids = append(ids, mg.Id)
			}
		}
	}

	if len(ids) == 0 {
		return nil, pgx.ErrNoRows
	}
	return ids, nil
}
