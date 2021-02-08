package handle

import (
	"gymrat/db/repo"
	"gymrat/types"
	"net/http"

	"github.com/bouk/httprouter"
)

// CreateWorkoutExercise handles POST requests.
func CreateWorkoutExercise(w http.ResponseWriter, r *http.Request) {
	var workoutExercise types.WorkoutExercise
	err := readBody(r, &workoutExercise)
	if err != nil {
		respondWithError(w, err)
		return
	}

	savedWorkoutExercise, err := repo.CreateWorkoutExercise(r.Context(), workoutExercise)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, savedWorkoutExercise, http.StatusCreated)
}

// DeleteWorkoutExercise handles DELETE requests.
func DeleteWorkoutExercise(w http.ResponseWriter, r *http.Request) {
	workoutId := httprouter.GetParam(r, "id")
  exerciseId := httprouter.GetParam(r, "exerciseId")

	err := repo.DeleteWorkoutExercise(r.Context(), workoutId, exerciseId)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithStatus(w, http.StatusOK)
}

// ListWorkoutExercises handles list GET requests.
func ListWorkoutExercises(w http.ResponseWriter, r *http.Request) {
  id := httprouter.GetParam(r, "id")

	workoutExercises, err := repo.ListWorkoutExercises(r.Context(), id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, workoutExercises, http.StatusOK)
}
