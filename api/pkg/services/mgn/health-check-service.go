package services

import (
	"brahmafi-build-it/api/pkg/database"
)

type HealthStatus string

const (
	Health   HealthStatus = "Health"
	Unhealth HealthStatus = "Unhealth"
)

type HealthResponse struct {
	Database HealthStatus
}

func HandleHealthCheck() (HealthResponse, error) {
	response := HealthResponse{}

	err := database.Ping()

	if err != nil {
		response.Database = Unhealth
	} else {
		response.Database = Health
	}

	return response, err
}
