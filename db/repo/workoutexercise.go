package repo

import (
	"context"
	"database/sql"
	"gymrat/db"
	"gymrat/types"
)

// CreateWorkoutExercise creates a workout_exercise in the DB.
func CreateWorkoutExercise(ctx context.Context, workoutExercise types.WorkoutExercise) (types.WorkoutExercise, error) {
	const insertWorkoutExercise = `
		INSERT INTO workout_exercise (
			workoutId,
      exerciseId
		) VALUES (
			:workoutId,
			:exerciseId
		)`

	_, err := db.ContextDB(ctx).NamedExec(insertWorkoutExercise, workoutExercise)

	if err != nil {
		return types.WorkoutExercise{}, err
	}

	return workoutExercise, nil
}

// DeleteWorkoutExercise deletes the workout_exercise..
// It will *not* return an error if no workout was found for that ID.
func DeleteWorkoutExercise(ctx context.Context, workoutId string, exerciseId string) error {
  const deleteWorkout = `
    DELETE
    FROM workout_exercise
    WHERE workoutId = ? AND exerciseId = ?`

  _, err := db.ContextDB(ctx).Exec(deleteWorkout, workoutId, exerciseId)
  return err
}

// ListWorkoutExercises
func ListWorkoutExercises(ctx context.Context, workoutId string) ([]types.Exercise, error) {
	const list = `
		SELECT id, name
		FROM workout_exercise
		INNER JOIN exercise on exercise.id = workout_exercise.exerciseId
    WHERE workoutId = ?`

	exercises := []types.Exercise{}
  err := db.ContextDB(ctx).Select(&exercises, list, workoutId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewMissingEntityError(workoutId)
		}
		return nil, err
	}

	return exercises, err
}
