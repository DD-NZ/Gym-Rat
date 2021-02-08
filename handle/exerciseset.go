package handle

import (
	"gymrat/db/repo"
	"gymrat/types"
	"net/http"

	"github.com/bouk/httprouter"
)

// GetExercise handles GET requests.
func GetExerciseSet(w http.ResponseWriter, r *http.Request) {
	id := httprouter.GetParam(r, "id")

	exerciseSet, err := repo.GetExerciseSetByID(r.Context(), id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, exerciseSet, http.StatusOK)
}

// CreateExercise handles POST requests.
func CreateExerciseSet(w http.ResponseWriter, r *http.Request) {
	var exerciseSet types.ExerciseSet
	err := readBody(r, &exerciseSet)
	if err != nil {
		respondWithError(w, err)
		return
	}

	savedExerciseSet, err := repo.CreateExerciseSet(r.Context(), exerciseSet)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, savedExerciseSet, http.StatusCreated)
}

// UpdateExercise handles PUT requests.
func UpdateExerciseSet(w http.ResponseWriter, r *http.Request) {
	var exerciseSet types.ExerciseSet
	err := readBody(r, &exerciseSet)
	if err != nil {
		respondWithError(w, err)
		return
	}

	id := httprouter.GetParam(r, "id")

	updatedExerciseSet, err := repo.UpdateExerciseSet(r.Context(), id, exerciseSet)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, updatedExerciseSet, http.StatusOK)
}

// DeleteExercise handles DELETE requests.
func DeleteExerciseSet(w http.ResponseWriter, r *http.Request) {
	id := httprouter.GetParam(r, "id")

	err := repo.DeleteExerciseSet(r.Context(), id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithStatus(w, http.StatusOK)
}

// ListExercsies handles list GET requests.
func ListExerciseSets(w http.ResponseWriter, r *http.Request) {
	id := httprouter.GetParam(r, "id")
	exerciseSets, err := repo.ListExerciseSets(r.Context(), id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, exerciseSets, http.StatusOK)
}
