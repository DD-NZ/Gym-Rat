package types

// Workout represents a workout in the DB and API.
type Workout struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
}
