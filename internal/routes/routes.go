package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/matiasalek/go-http/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutByID)

	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
	r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkoutById)
	r.Delete("/workouts/{id}", app.WorkoutHandler.HandleDeleteWorkoutById)

	r.Post("/users", app.UserHandler.HandleRegisterUser)

	return r
}
