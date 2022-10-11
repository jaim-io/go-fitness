package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ExerciseMuscleGroupsContext struct {
	DB *pgxpool.Pool
}

func (c ExerciseMuscleGroupsContext) RemoveAllByExercise(exercise Exercise) error {
	_, err := c.DB.Exec(context.Background(), `
		DELETE FROM exercise_muscle_groups as emg
		WHERE emg.exercise_id = $1
	`, exercise.Id)

	if err != nil {
		return err
	}

	return nil
}

func (c ExerciseMuscleGroupsContext) RemoveAllByMuscleGroup(muscleGroup MuscleGroup) error {
	_, err := c.DB.Exec(context.Background(), `
		DELETE FROM exercise_muscle_groups as emg
		WHERE emg.muscle_group_id = $1
	`, muscleGroup.Id)

	if err != nil {
		return err
	}

	return nil
}

func (c ExerciseMuscleGroupsContext) SetFromExercise(exerciseId uint32, muscleGroupIds []uint32) error {
	var insertValues string
	for i, id := range muscleGroupIds {
		// Example: (1, 3), (1, 2)
		insertValues += fmt.Sprintf("(%d, %d)", exerciseId, id)
		if i != len(muscleGroupIds)-1 {
			insertValues += ", "
		}
	}

	query := fmt.Sprintf(`
		INSERT INTO exercise_muscle_groups (exercise_id, muscle_group_id)
		VALUES %s
	`, insertValues)
	_, err := c.DB.Exec(context.Background(), query)

	if err != nil {
		return err
	}

	return nil
}
