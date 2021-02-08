package server

import (
	"gymrat/db"
	"gymrat/handle"
	"net/http"

	"github.com/bouk/httprouter"
)

func New() http.Handler {
	router := httprouter.New()

	router.GET("/", handle.HomeResponse)

	router.GET("/api/workout/:id", handle.GetWorkout)
	router.POST("/api/workout", handle.CreateWorkout)
	router.PUT("/api/workout/:id", handle.UpdateWorkout)
	router.DELETE("/api/workout/:id", handle.DeleteWorkout)
	router.GET("/api/workouts", handle.ListWorkouts)

	router.GET("/api/exercise/:id", handle.GetExercise)
	router.POST("/api/exercise", handle.CreateExercise)
	router.PUT("/api/exercise/:id", handle.UpdateExercise)
	router.DELETE("/api/exercise/:id", handle.DeleteExercise)
	router.GET("/api/exercises", handle.ListExercises)

	router.POST("/api/workoutexercise", handle.CreateWorkoutExercise)
	router.GET("/api/workoutexercise/:id", handle.ListWorkoutExercises)
	router.DELETE("/api/workout/:id/exercise/:exerciseId", handle.DeleteWorkoutExercise)

	router.GET("/api/day/:id", handle.GetDay)
	router.POST("/api/day", handle.CreateDay)
	router.DELETE("/api/day/:id", handle.DeleteDay)
	router.GET("/api/days/:id", handle.ListDays)

	router.GET("/api/exerciseset/:id", handle.GetExerciseSet)
	router.POST("/api/exerciseset", handle.CreateExerciseSet)
	router.PUT("/api/exerciseset/:id", handle.UpdateExerciseSet)
	router.DELETE("/api/exerciseset/:id", handle.DeleteExerciseSet)
	router.GET("/api/exercisesets/:id", handle.ListExerciseSets)

	return &Server{
		router: router,
	}
}

type Server struct {
	router http.Handler
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = db.ContextWithDB(ctx, db.Client())
	server.router.ServeHTTP(w, r.WithContext(ctx))
}
