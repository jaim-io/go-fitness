package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Exercise struct {
	Id           uint32   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	MuscleGroups []string `json:"muscle_groups"`
	ImagePath    string   `json:"image_path"`
	VideoLink    string   `json:"video_link"`
}

// Create a custom ExerciseContext type which wraps the sql.DB connection pool.
type ExerciseContext struct {
	DB *pgxpool.Pool
}

func (c ExerciseContext) GetAll() ([]Exercise, error) {
	rows, err := c.DB.Query(context.Background(), `
		SELECT 
			e.id, 
			e.name, 
			e.description, 
			ARRAY(
				SELECT mg.name
				FROM muscle_groups as mg 
				JOIN exercise_muscle_groups as emg
					ON mg.id=emg.muscle_group_id
				WHERE emg.exercise_id=e.id) as "muscle_groups",
			e.image_path, 
			e.video_link 
		FROM exercises as e
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exs []Exercise

	for rows.Next() {
		var ex Exercise

		if err := rows.Scan(&ex.Id, &ex.Name, &ex.Description, &ex.MuscleGroups, &ex.ImagePath, &ex.VideoLink); err != nil {
			return nil, err
		}
		exs = append(exs, ex)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exs, nil
}

func (c ExerciseContext) GetById(id uint32) (Exercise, error) {
	var ex Exercise

	row := c.DB.QueryRow(context.Background(), `
		SELECT 
			id, 
			name, 
			description, 
			ARRAY(
				SELECT DISTINCT ON (mg.name) mg.name 
				FROM muscle_groups as mg 
				JOIN exercise_muscle_groups as emg
					ON emg.muscle_group_id=mg.id
				WHERE emg.exercise_id=$1) as "muscle_groups",
			image_path, 
			video_link 
		FROM exercises 
		WHERE id=$1;
	`, id)
	if err := row.Scan(&ex.Id, &ex.Name, &ex.Description, &ex.MuscleGroups, &ex.ImagePath, &ex.VideoLink); err != nil {
		return ex, err
	}
	return ex, nil
}

func (c ExerciseContext) Update(id uint32, exercise Exercise) error {
	_, err := c.DB.Exec(context.Background(), `
		UPDATE exercises
		SET name=$1, SET description=$2, SET image_path=$3, SET video_link=$4
		WHERE id=$5
		`, exercise.Name, exercise.Description, exercise.ImagePath, exercise.VideoLink, exercise.Id,
	)

	if err != nil {
		return err
	}
	return nil
}

func (c ExerciseContext) Add(exercise Exercise) (Exercise, error) {
	_, err := c.DB.Exec(context.Background(), `
		INSERT INTO exercises (name, description, image_path, video_link)
		VALUES ($1, $2, $3, $4)`, exercise.Name, exercise.Description, exercise.ImagePath, exercise.VideoLink,
	)

	if err != nil {
		return Exercise{}, err
	}
	return exercise, nil
}

func (c ExerciseContext) Remove(exercise Exercise) error {
	_, err := c.DB.Exec(context.Background(), `
		DELETE FROM exercises 
		WHERE id=$1 AND name=$2 AND description=$3 AND image_path=$4, AND video_link=$5`,
		exercise.Id, exercise.Name, exercise.Description, exercise.ImagePath, exercise.VideoLink,
	)
	if err != nil {
		return err
	}
	return nil
}
