package main

import (
	//...

	"applicationDesignTest/cmd/booking"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	app := booking.InitApp(context.Background())

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/orders", app.Handlers.Booking.CreateOrders)
	http.ListenAndServe(":8080", r)
}
