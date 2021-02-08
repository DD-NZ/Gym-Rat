package types

// WorkoutExercise represents a workout_exercise in the DB and API.
type WorkoutExercise struct {
  WorkoutId       string `json:"workoutId"`
  ExerciseId      string `json:"exerciseId"`
}
