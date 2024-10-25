package mocks

import (
	"context"
	"database/sql"
	"github.com/abdukhashimov/go_api_mono_repo/internal/core/repository/psql/sqlc"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testQueries *sqlc.Queries
	testDB      *sql.DB
)

// initializeDB establishes the database connection and applies migrations.
func initializeDB(psqlUri string) (*pgxpool.Pool, error) {
	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", psqlUri)
	if err != nil {
		return nil, err
	}

	// Apply migrations
	migrationPath := filepath.Join("..", "migrate", "migrations")
	if err := goose.Up(db, migrationPath); err != nil {
		return nil, err
	}

	// Initialize pgx pool for queries
	dbConn, err := pgxpool.Connect(context.Background(), psqlUri)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

// cleanUpDB truncates the tables and resets the identity.
func cleanUpDB(db *sql.DB) error {
	_, err := db.Exec(`TRUNCATE TABLE "user", "todo" RESTART IDENTITY CASCADE;`)
	return err
}

func TestMain(m *testing.M) {
	psqlUri := "postgresql://postgres:postgres@localhost:5432/testdb?sslmode=disable"

	// Initialize the database and handle errors
	var err error
	testDB, err = sql.Open("postgres", psqlUri)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	defer testDB.Close()

	dbConn, err := initializeDB(psqlUri)
	if err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}
	defer dbConn.Close() // Ensure the pool is closed after tests

	// Initialize Queries
	testQueries = sqlc.New(dbConn)

	// Run tests
	code := m.Run()

	// Clean up database after tests
	if err := cleanUpDB(testDB); err != nil {
		log.Fatalf("cannot clean up db: %v", err)
	}

	os.Exit(code)
}
