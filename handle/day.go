package handle

import (
	"gymrat/db/repo"
	"gymrat/types"
	"net/http"

	"github.com/bouk/httprouter"
)

// GetDay handles GET requests.
func GetDay(w http.ResponseWriter, r *http.Request) {
	id := httprouter.GetParam(r, "id")

	day, err := repo.GetDayByID(r.Context(), id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, day, http.StatusOK)
}

// CreateWorkout handles POST requests.
func CreateDay(w http.ResponseWriter, r *http.Request) {
	var day types.Day
	err := readBody(r, &day)
	if err != nil {
		respondWithError(w, err)
		return
	}

	savedDay, err := repo.CreateDay(r.Context(), day)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, savedDay, http.StatusCreated)
}

// DeleteWorkout handles DELETE requests.
func DeleteDay(w http.ResponseWriter, r *http.Request) {
	id := httprouter.GetParam(r, "id")

	err := repo.DeleteDay(r.Context(), id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithStatus(w, http.StatusOK)
}

// ListWorkouts handles list GET requests.
func ListDays(w http.ResponseWriter, r *http.Request) {
  id := httprouter.GetParam(r, "id")
	days, err := repo.ListDays(r.Context(), id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, days, http.StatusOK)
}
