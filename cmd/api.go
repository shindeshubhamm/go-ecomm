package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repository "github.com/shindeshubhamm/go-ecomm/internal/adapters/postgresql/sqlc"
	"github.com/shindeshubhamm/go-ecomm/internal/service"
	"github.com/shindeshubhamm/go-ecomm/internal/transport/http/handlers"
)

type application struct {
	config config
	db     *pgx.Conn
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)                 // important for rate limiting
	r.Use(middleware.RealIP)                    // important for rate limiting, analytics and tracing
	r.Use(middleware.Logger)                    // important for logging
	r.Use(middleware.Recoverer)                 // important for recovery
	r.Use(middleware.Timeout(60 * time.Second)) // important for request timeout

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good."))
	})

	repo := repository.New(app.db)
	r.Route("/products", func(r chi.Router) {
		productService := service.NewProductService(repo)
		productHandler := handlers.NewProductHandler(productService)
		r.Get("/", productHandler.ListProducts)
	})

	// r.Route("/orders", func(r chi.Router) {
	// 	orderHandler := handlers.NewOrderHandler(nil)
	// 	r.Get("/", orderHandler.ListOrders)
	// })

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("> Server running on %s", app.config.addr)
	return srv.ListenAndServe()
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
