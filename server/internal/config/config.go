package config

import (
	"log"
	"os"
	"sync"

	"assert200.com/vincivia/internal/logger"
)

// Config struct
type Config struct {
	Port               string
	ServerHost         string
	GoogleClientID     string
	GoogleClientSecret string
	EncryptionKey      string
	DBDriver           string
	DBUsername         string
	DBPassword         string
	DBHost             string
	DBName             string
}

var once sync.Once
var instance Config

// Get the config instance
func Get() Config {
	once.Do(func() { // <-- atomic, does not allow repeating

		port, found := os.LookupEnv("PORT")
		if !found {
			logger.Get().Println("WARNING: Evironment variable PORT was not found, defaulting to 8080")
			port = "8080"
		}

		serverHost := mustGetenv("REACT_APP_SERVER_HOST")
		googleClientID := mustGetenv("REACT_APP_GOOGLE_CLIENT_ID")
		googleClientSecret := mustGetenv("GOOGLE_CLIENT_SECRET")
		encryptionKey := mustGetenv("ENCRYPTION_KEY")
		dbUsername := mustGetenv("DB_USERNAME")
		dbPassword := mustGetenv("DB_PASSWORD")
		dbHost := mustGetenv("DB_HOST")
		dbName := mustGetenv("DB_NAME")
		dbDriver := mustGetenv("DB_DRIVER")

		config := Config{
			Port:               port,
			ServerHost:         serverHost,
			GoogleClientID:     googleClientID,
			GoogleClientSecret: googleClientSecret,
			EncryptionKey:      encryptionKey,
			DBUsername:         dbUsername,
			DBPassword:         dbPassword,
			DBHost:             dbHost,
			DBName:             dbName,
			DBDriver:           dbDriver,
		}

		instance = config

		logger.Get().Println("Config singleton instantiated")
	})

	return instance
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
