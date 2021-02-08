package repo

import (
	"context"
	"database/sql"
	"gymrat/db"
	"gymrat/types"

	"github.com/gofrs/uuid"
)

// CreateExercise creates an exercise in the DB.
// It will generate the ID itself.
func CreateExercise(ctx context.Context, exercise types.Exercise) (types.Exercise, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return types.Exercise{}, err
	}

	exercise.ID = id.String()

	const insertExercise = `
		INSERT INTO exercise (
			id,
			name
		) VALUES (
			:id,
			:name
		)`

	_, err = db.ContextDB(ctx).NamedExec(insertExercise, exercise)

	if err != nil {
		return types.Exercise{}, err
	}

	return exercise, nil
}

// GetExerciseByID returns the exercise associated with the provided ID.
// It will return nil (and no error) if no workout was found.
func GetExerciseByID(ctx context.Context, id string) (*types.Exercise, error) {
	const getExercise = `
		SELECT id, name
		FROM exercise
		WHERE id = ?`

	var exercise types.Exercise
	err := db.ContextDB(ctx).Get(&exercise, getExercise, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewMissingEntityError(id)
		}
		return nil, err
	}

	return &exercise, err
}

// UpdateExercise updates the exercise associated with the provided ID
// with the provided product value. It will return nil if there was no
// workout for that ID.
func UpdateExercise(ctx context.Context, id string, exercise types.Exercise) (*types.Exercise, error) {
	const updateExercise = `
		UPDATE exercise SET
			name = :name
		WHERE id = :id`

	exercise.ID = id

	result, err := db.ContextDB(ctx).NamedExec(updateExercise, exercise)
	if err != nil {
		return nil, err
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAff == 0 {
		return nil, types.NewMissingEntityError(id)
	}

	return &exercise, nil
}

// DeleteExercise deletes the exercise associated with the provided ID.
// It will *not* return an error if no workout was found for that ID.
func DeleteExercise(ctx context.Context, id string) error {
	tx, err := db.ContextDB(ctx).Beginx()

	if (err != nil) {
		return err
	}

	const deleteExerciseFromWorkouts = `
			Delete
			From workout_exercise
			where exerciseId = ?
	`;

	_, err = tx.Exec(deleteExerciseFromWorkouts, id)

	if (err != nil) {
		tx.Rollback()
		return err
	}

	const deleteExercise = `
		DELETE
		FROM exercise
		WHERE id = ?`

	_, err = tx.Exec(deleteExercise, id);

	if (err != nil) {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	if (err != nil) {
		tx.Rollback()
		return err
	}

	return err
}

// ListExercises lists exercises in the DB.
func ListExercises(ctx context.Context) ([]types.Exercise, error) {
	const list = `
		SELECT id, name
		FROM exercise`

	exercises := []types.Exercise{}
	err := db.ContextDB(ctx).Select(&exercises, list)
	return exercises, err
}
