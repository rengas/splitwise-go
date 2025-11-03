package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	database "splitwise-go/internal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting application...")

	// 1. Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set in .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// 2. Connect to Postgres database
	dbPool, err := database.NewConnection(databaseURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	// Defer closing the connection pool
	defer dbPool.Close()
	log.Println("Database connection successful.")

	// 3. Set up chi router
	r := chi.NewRouter()

	// Add middleware
	r.Use(middleware.Logger)    // Log requests
	r.Use(middleware.Recoverer) // Recover from panics

	// 4. Healthcheck route
	// We pass the dbPool to the handler
	r.Get("/health", healthCheckHandler(dbPool))

	// 5. Start the server
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

// healthCheckHandler is a handler function that pings the database.
func healthCheckHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 4a. Query the database with "SELECT 1"
		var result int
		err := db.QueryRow(r.Context(), "SELECT 1").Scan(&result)

		if err != nil {
			log.Printf("Health check failed (db query): %v", err)
			http.Error(w, "Database connection failed", http.StatusInternalServerError)
			return
		}

		if result != 1 {
			log.Printf("Health check failed (unexpected result): %d", result)
			http.Error(w, "Database query failed", http.StatusInternalServerError)
			return
		}

		// 4b. Respond with a success message
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Using Fprintln to write to the ResponseWriter
		fmt.Fprintln(w, `{"status": "ok", "message": "I am healthy!"}`)
	}
}
