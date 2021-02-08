package handle

import (
	"gymrat/db/repo"
	"gymrat/types"
	"net/http"

	"github.com/bouk/httprouter"
)

// GetWorkout handles GET requests.
func GetWorkout(w http.ResponseWriter, r *http.Request) {
	id := httprouter.GetParam(r, "id")

	workout, err := repo.GetWorkoutByID(r.Context(), id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, workout, http.StatusOK)
}

// CreateWorkout handles POST requests.
func CreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout types.Workout
	err := readBody(r, &workout)
	if err != nil {
		respondWithError(w, err)
		return
	}

	savedWorkout, err := repo.CreateWorkout(r.Context(), workout)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, savedWorkout, http.StatusCreated)
}

// UpdateWorkout handles PUT requests.
func UpdateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout types.Workout
	err := readBody(r, &workout)
	if err != nil {
		respondWithError(w, err)
		return
	}

	id := httprouter.GetParam(r, "id")

	updatedWorkout, err := repo.UpdateWorkout(r.Context(), id, workout)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, updatedWorkout, http.StatusOK)
}

// DeleteWorkout handles DELETE requests.
func DeleteWorkout(w http.ResponseWriter, r *http.Request) {
	id := httprouter.GetParam(r, "id")

	err := repo.DeleteWorkout(r.Context(), id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithStatus(w, http.StatusOK)
}

// ListWorkouts handles list GET requests.
func ListWorkouts(w http.ResponseWriter, r *http.Request) {
	workouts, err := repo.ListWorkouts(r.Context())
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, workouts, http.StatusOK)
}
