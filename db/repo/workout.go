package repo

import (
	"context"
	"database/sql"
	"gymrat/db"
	"gymrat/types"

	"github.com/gofrs/uuid"
)

// CreateWorkout creates a workout in the DB.
// It will generate the ID itself.
func CreateWorkout(ctx context.Context, workout types.Workout) (types.Workout, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return types.Workout{}, err
	}

	workout.ID = id.String()

	const insertWorkout = `
		INSERT INTO workout (
			id,
			name
		) VALUES (
			:id,
			:name
		)`

	_, err = db.ContextDB(ctx).NamedExec(insertWorkout, workout)

	if err != nil {
		return types.Workout{}, err
	}

	return workout, nil
}

// GetWorkoutByID returns the workout associated with the provided ID.
// It will return nil (and no error) if no workout was found.
func GetWorkoutByID(ctx context.Context, id string) (*types.Workout, error) {
	const getWorkout = `
		SELECT id, name
		FROM workout
		WHERE id = ?`

	var workout types.Workout
	err := db.ContextDB(ctx).Get(&workout, getWorkout, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewMissingEntityError(id)
		}
		return nil, err
	}

	return &workout, err
}

// UpdateWorkout updates the product associated with the provided ID with the provided product value.
// It will return nil if there was no workout for that ID.
func UpdateWorkout(ctx context.Context, id string, workout types.Workout) (*types.Workout, error) {
	const updateWorkout = `
		UPDATE workout SET
			name = :name
		WHERE id = :id`

	workout.ID = id

	result, err := db.ContextDB(ctx).NamedExec(updateWorkout, workout)
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

	return &workout, nil
}

// DeleteWorkout deletes the workout associated with the provided ID.
// It will *not* return an error if no workout was found for that ID.
func DeleteWorkout(ctx context.Context, id string) error {
	const deleteWorkout = `
		DELETE
		FROM workout
		WHERE id = ?`

	_, err := db.ContextDB(ctx).Exec(deleteWorkout, id)
	return err
}

// ListWorkouts lists workouts in the DB.
func ListWorkouts(ctx context.Context) ([]types.Workout, error) {
	const list = `
		SELECT id, name
		FROM workout`

	workouts := []types.Workout{}
	err := db.ContextDB(ctx).Select(&workouts, list)
	return workouts, err
}
