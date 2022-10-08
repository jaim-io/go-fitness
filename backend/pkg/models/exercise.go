package models

import (
	"database/sql"
)

type Exercise struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

// Create a custom ExerciseContext type which wraps the sql.DB connection pool.
type ExerciseContext struct {
	DB *sql.DB
}

func (c ExerciseContext) GetAll() ([]Exercise, error) {
	rows, err := c.DB.Query("SELECT * FROM exercises")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exs []Exercise

	for rows.Next() {
		var ex Exercise

		if err := rows.Scan(&ex.Id, &ex.Name); err != nil {
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

	row := c.DB.QueryRow("SELECT * FROM exercises WHERE id = ?", id)
	if err := row.Scan(&ex.Id, &ex.Name); err != nil {
		return ex, err
	}
	return ex, nil
}

func (c ExerciseContext) Update(id uint32, exercise Exercise) error {
	_, err := c.DB.Exec(`
		UPDATE exercises
		SET name = ? 
		WHERE id = ?
		`, exercise.Name, exercise.Id,
	)

	if err != nil {
		return err
	}
	return nil
}

func (c ExerciseContext) Add(exercise Exercise) (Exercise, error) {
	_, err := c.DB.Exec(`
		INSERT INTO exercises (name)
		VALUES (?)`, exercise.Name,
	)

	if err != nil {
		return Exercise{}, err
	}
	return exercise, nil
}

func (c ExerciseContext) Remove(exercise Exercise) error {
	_, err := c.DB.Exec(`
		DELETE FROM exercises 
		WHERE id = ? AND name = ?`,
		exercise.Id, exercise.Name,
	)
	if err != nil {
		return err
	}
	return nil
}
