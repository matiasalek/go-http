package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/matiasalek/go-http/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)

		r.Get("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.HandleGetWorkoutByID))
		r.Post("/workouts", app.Middleware.RequireUser(app.WorkoutHandler.HandleCreateWorkout))
		r.Put("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.HandleUpdateWorkoutByID))
		r.Delete("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.HandleDeleteWorkoutByID))
	})

	r.Post("/users", app.UserHandler.HandleRegisterUser)
	r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)

	return r
}
