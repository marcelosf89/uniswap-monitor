package configurations

import (
	"os"
)

const (
	DEFAULT_PORT = "3000"
)

func GetDatabaseConnectionString() string {
	return os.Getenv("DATABASE_CONNECTION_STRING")
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return DEFAULT_PORT
	}

	return port
}
