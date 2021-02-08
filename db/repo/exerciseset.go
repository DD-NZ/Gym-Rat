package repo

import (
	"context"
	"database/sql"
	"gymrat/db"
	"gymrat/types"

	"github.com/gofrs/uuid"
)

// CreateExerciseSet creates an exerciseset in the DB.
// It will generate the ID itself.
func CreateExerciseSet(ctx context.Context, exerciseSet types.ExerciseSet) (types.ExerciseSet, error) {
  id, err := uuid.NewV4()
	if err != nil {
		return types.ExerciseSet{}, err
	}

	exerciseSet.ID = id.String()

  lastExerciseSet := types.ExerciseSet{}

  const findNextSetNo = `
    SELECT dayId, setNo, rep, weight
    FROM exerciseset
    WHERE dayId = ?
    ORDER BY setNo Desc
    Limit 1`

	err = db.ContextDB(ctx).Get(&lastExerciseSet, findNextSetNo, exerciseSet.DayId)
  if err != nil {
    if err != sql.ErrNoRows {
        return types.ExerciseSet{}, err
    } else {
      exerciseSet.SetNo = 1
    }
  } else {
    exerciseSet.SetNo = lastExerciseSet.SetNo + 1
  }

	const insertExerciseSet = `
		INSERT INTO exerciseset (
			id, dayId, setNo, rep, weight
		) VALUES (
			:id, :dayId, :setNo, :rep, :weight
		)`

	_, err = db.ContextDB(ctx).NamedExec(insertExerciseSet, exerciseSet)

	if err != nil {
		return types.ExerciseSet{}, err
	}

	return exerciseSet, nil
}

// GetExerciseSetByID returns the exerciseset associated with the provided ID.
// It will return nil (and no error) if no workout was found.
func GetExerciseSetByID(ctx context.Context, id string) (*types.ExerciseSet, error) {
	const getExerciseSet = `
		SELECT id, dayId, setNo, rep, weight
		FROM exerciseset
		WHERE id = ?`

	var exerciseSet types.ExerciseSet
	err := db.ContextDB(ctx).Get(&exerciseSet, getExerciseSet, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewMissingEntityError(id)
		}
		return nil, err
	}

	return &exerciseSet, err
}

// UpdateExerciseSet updates the exerciseset associated with the provided ID.
// It will return nil if there was no workout for that ID.
func UpdateExerciseSet(ctx context.Context, id string, exerciseSet types.ExerciseSet) (*types.ExerciseSet, error) {
	const updateExerciseSet = `
		UPDATE exerciseset SET
			rep = :rep,
      weight = :weight
		WHERE id = :id`

	exerciseSet.ID = id

	result, err := db.ContextDB(ctx).NamedExec(updateExerciseSet, exerciseSet)
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

	return &exerciseSet, nil
}

// DeleteExerciseSet deletes the exerciseset associated with the provided ID.
// It will *not* return an error if no workout was found for that ID.
func DeleteExerciseSet(ctx context.Context, id string) error {
	exerciseSet, err := GetExerciseSetByID(ctx, id)
	if err != nil {
		return err
	}

	tx, err := db.ContextDB(ctx).Beginx()

	if (err != nil) {
		return err
	}

	const deleteExerciseSet =`
		DELETE
		FROM exerciseset
		WHERE id = :id`

	result, err := tx.NamedExec(deleteExerciseSet, exerciseSet)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAff == 0 {
		tx.Rollback()
		return types.NewMissingEntityError(id)
	}

	const updateExerciseSetSetNo =`
		UPDATE exerciseset SET
			setNo = setNo - 1
		WHERE dayId = :dayId AND setNo > :setNo
	`;

	result, err = tx.NamedExec(updateExerciseSetSetNo, exerciseSet)
	if err != nil {
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

// ListExerciseSets lists exercisesets in the DB with the provided dayId.
func ListExerciseSets(ctx context.Context, dayId string) ([]types.ExerciseSet, error) {
	const list = `
		SELECT id, dayId, setNo, rep, weight
		FROM exerciseset
		WHERE dayId = ?
		ORDER BY setNo ASC`

	exerciseSets := []types.ExerciseSet{}
	err := db.ContextDB(ctx).Select(&exerciseSets, list, dayId)
	return exerciseSets, err
}
