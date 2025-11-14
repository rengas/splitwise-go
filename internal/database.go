package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewConnection creates and returns a new pgxpool.Pool connection.
// It also pings the database to verify the connection is live.
func NewConnection(databaseURL string) (*pgxpool.Pool, error) {
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL cannot be empty")
	}

	fmt.Println(databaseURL)

	// Parse the connection string
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	// Set some sensible defaults for the connection pool
	config.MaxConns = 10                      // Max number of connections
	config.MinConns = 2                       // Min number of connections
	config.MaxConnIdleTime = time.Minute * 30 // Duration before an idle connection is closed
	config.MaxConnLifetime = time.Hour        // Max duration a connection can exist
	config.HealthCheckPeriod = time.Minute    // Frequency of health checks

	// Attempt to connect to the database
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Ping the database to ensure connectivity
	// We use a context with a timeout for the ping
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = pool.Ping(ctx); err != nil {
		pool.Close() // Close the pool if the initial ping fails
		return nil, fmt.Errorf("failed to ping database on connect: %w", err)
	}

	log.Println("Database connection pool established.")
	return pool, nil
}
