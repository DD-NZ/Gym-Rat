package types

// Exercise represents a exercise in the DB and API.
type Day struct {
  ID         string `json:"id"`
  DayDate    string `json:"dayDate"`
  ExerciseId string `json:"exerciseId"`
}
