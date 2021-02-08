package repo

import (
  "time"
	"context"
	"database/sql"
	"gymrat/db"
	"gymrat/types"

	"github.com/gofrs/uuid"
)

// CreateDay creates a day in the DB.
// It will generate the ID, DateTime itself.
func CreateDay(ctx context.Context, day types.Day) (types.Day, error) {
	id, err := uuid.NewV4()

	if err != nil {
		return types.Day{}, err
	}

  loc, _ := time.LoadLocation("Pacific/Auckland")
  day.ID = id.String()
  day.DayDate = time.Now().In(loc).Format("2006-01-02 15:04:05")

	const insertExercise = `
		INSERT INTO day (
			id,
			dayDate,
      exerciseId
		) VALUES (
			:id,
	    :dayDate,
      :exerciseId
		)`

	_, err = db.ContextDB(ctx).NamedExec(insertExercise, day)

	if err != nil {
		return types.Day{}, err
	}

	return day, nil
}

// GetDayByID returns the day associated with the provided ID.
// It will return nil (and no error) if no day was found.
func GetDayByID(ctx context.Context, dayId string) (*types.Day, error) {
	const getDay = `
		SELECT id, dayDate, exerciseId
		FROM day
		WHERE id = ?`

	var day types.Day
	err := db.ContextDB(ctx).Get(&day, getDay, dayId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewMissingEntityError(dayId)
		}
		return nil, err
	}

	return &day, err
}

// DeleteDay deletes the day associated with the provided ID.
// It will *not* return an error if no workout was found for that ID.
func DeleteDay(ctx context.Context, dayId string) error {
  const deleteDay = `
		DELETE
		FROM day
		WHERE id = ?`

	_, err := db.ContextDB(ctx).Exec(deleteDay, dayId)
	return err
}

// ListDay lists days in the DB.
func ListDays(ctx context.Context, exerciseId string) ([]types.Day, error) {
	const list = `
		SELECT id, dayDate, exerciseId
		FROM day
    WHERE exerciseId = ?`

	days := []types.Day{}
	err := db.ContextDB(ctx).Select(&days, list, exerciseId)
	return days, err
}
