package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"assert200.com/vincivia/internal/config"
	"assert200.com/vincivia/internal/logger"

	// Google Public SQL required this poxy
	// So we can sql.Open("cloudsqlpostgres")
	// https://cloud.google.com/sql/docs/postgres/connect-external-app#go
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
)

var once sync.Once
var instance *sql.DB

//Get the db connection instance
func Get() *sql.DB {
	once.Do(func() { // <-- atomic, does not allow repeating
		db, err := openDB()
		if err != nil {
			logger.Get().Fatal("Database connection failed with: ", err)
		}
		instance = db

		logger.Get().Println("Database connection singleton instantiated")
	})

	return instance
}

func openDB() (*sql.DB, error) {
	config := config.Get()
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		config.DBHost,
		config.DBName,
		config.DBUsername,
		config.DBPassword)
	dbPool, err := sql.Open(config.DBDriver, dsn)

	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	return dbPool, nil
}

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}
