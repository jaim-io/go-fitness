package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MuscleGroup struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
}

// Create a custom ExerciseContext type which wraps the sql.DB connection pool.
type MuscleGroupContext struct {
	DB *pgxpool.Pool
}

func (c MuscleGroupContext) GetAll() ([]MuscleGroup, error) {
	rows, err := c.DB.Query(context.Background(), `
		SELECT 
			*
		FROM muscle_groups
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mgs []MuscleGroup

	for rows.Next() {
		var mg MuscleGroup

		if err := rows.Scan(&mg.Id, &mg.Name, &mg.Description, &mg.ImagePath); err != nil {
			return nil, err
		}
		mgs = append(mgs, mg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return mgs, nil
}

func (c MuscleGroupContext) GetById(id uint32) (MuscleGroup, error) {
	var mg MuscleGroup

	row := c.DB.QueryRow(context.Background(), `
		SELECT 
			*
		FROM muscle_groups
		WHERE id=$1;
	`, id)
	if err := row.Scan(&mg.Id, &mg.Name, &mg.Description, &mg.ImagePath); err != nil {
		return mg, err
	}
	return mg, nil
}

func (c MuscleGroupContext) Update(id uint32, muscleGroup MuscleGroup) error {
	_, err := c.DB.Exec(context.Background(), `
		UPDATE muscle_groups
		SET name=$1, SET description=$2, SET image_path=$3,
		WHERE id=$4
		`, muscleGroup.Name, muscleGroup.Description, muscleGroup.ImagePath, muscleGroup.Id,
	)

	if err != nil {
		return err
	}
	return nil
}

func (c MuscleGroupContext) Add(muscleGroup MuscleGroup) (MuscleGroup, error) {
	_, err := c.DB.Exec(context.Background(), `
		INSERT INTO muscle_groups (name, description, image_path)
		VALUES ($1, $2, $3)`, muscleGroup.Name, muscleGroup.Description, muscleGroup.ImagePath,
	)

	if err != nil {
		return MuscleGroup{}, err
	}
	return muscleGroup, nil
}

func (c MuscleGroupContext) Remove(muscleGroup MuscleGroup) error {
	_, err := c.DB.Exec(context.Background(), `
		DELETE FROM muscle_groups
		WHERE id=$1 AND name=$2 AND description=$3 AND image_path=$4`,
		muscleGroup.Id, muscleGroup.Name, muscleGroup.Description, muscleGroup.ImagePath,
	)
	if err != nil {
		return err
	}
	return nil
}