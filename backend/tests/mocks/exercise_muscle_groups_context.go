package mocks

import "github.com/Jaim010/jaim-io/backend/pkg/models"

type MockExerciseMuscleGroupsContext struct{}

func (c MockExerciseMuscleGroupsContext) RemoveAllByExercise(exercise models.Exercise) error {
	return nil
}
func (c MockExerciseMuscleGroupsContext) RemoveAllByMuscleGroup(muscleGroup models.MuscleGroup) error {
	return nil
}
func (c *MockExerciseMuscleGroupsContext) SetFromExercise(exerciseId uint32, muscleGroupIds []uint32) error {
	return nil
}
