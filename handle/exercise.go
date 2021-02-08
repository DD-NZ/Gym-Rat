package handle

import (
	"gymrat/db/repo"
	"gymrat/types"
	"net/http"

	"github.com/bouk/httprouter"
)

// GetExercise handles GET requests.
func GetExercise(w http.ResponseWriter, r *http.Request) {
	id := httprouter.GetParam(r, "id")

	exercise, err := repo.GetExerciseByID(r.Context(), id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, exercise, http.StatusOK)
}

// CreateExercise handles POST requests.
func CreateExercise(w http.ResponseWriter, r *http.Request) {
	var exercise types.Exercise
	err := readBody(r, &exercise)
	if err != nil {
		respondWithError(w, err)
		return
	}

	savedExercise, err := repo.CreateExercise(r.Context(), exercise)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, savedExercise, http.StatusCreated)
}

// UpdateExercise handles PUT requests.
func UpdateExercise(w http.ResponseWriter, r *http.Request) {
	var exercise types.Exercise
	err := readBody(r, &exercise)
	if err != nil {
		respondWithError(w, err)
		return
	}

	id := httprouter.GetParam(r, "id")

	updatedExercise, err := repo.UpdateExercise(r.Context(), id, exercise)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, updatedExercise, http.StatusOK)
}

// DeleteExercise handles DELETE requests.
func DeleteExercise(w http.ResponseWriter, r *http.Request) {
	id := httprouter.GetParam(r, "id")

	err := repo.DeleteExercise(r.Context(), id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithStatus(w, http.StatusOK)
}

// ListExercsies handles list GET requests.
func ListExercises(w http.ResponseWriter, r *http.Request) {
	exercises, err := repo.ListExercises(r.Context())
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, exercises, http.StatusOK)
}
