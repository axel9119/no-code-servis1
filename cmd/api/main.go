package main

import (
	"log"
	"net/http"

	"nocode-backend/db"
	"nocode-backend/handlers"
	"nocode-backend/middleware"

	"github.com/go-chi/chi/v5"
)

func main() {
	dsn := "postgres://axel@localhost:5432/nocode?sslmode=disable"
	db.Connect(dsn)

	r := chi.NewRouter()

	r.Post("/auth/register", handlers.Register)
	r.Post("/auth/login", handlers.Login)

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware) // JWT middleware
		r.Get("/companies", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Это защищённый маршрут /api/companies"))
		})
	})

	log.Println("✅ Backend ready on :8080")
	http.ListenAndServe(":8080", r)
}
