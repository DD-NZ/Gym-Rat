package types

// ExerciseSet represents a exerciseset in the DB and API.
type ExerciseSet struct {
	ID     string  `json:"id"`
  DayId  string  `json:"dayId"`
  SetNo  int     `json:"setNo"`
  Rep    int     `json:"rep"`
  Weight float32 `json:"weight"`
}
