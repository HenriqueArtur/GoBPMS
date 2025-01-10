// Package dependencies provides a centralized mechanism for constructing and managing
// application dependencies, such as the database connection and environment variables.
// This package is designed to simplify dependency injection and resource initialization
// within the application.
package dependencies

import (
	Infrastructure "github.com/HenriqueArtur/ProcessInGo/src/Infrastructure"
	"github.com/HenriqueArtur/ProcessInGo/src/Infrastructure/database"
)

// Dependency is a container struct that holds essential components required
// across the application. It centralizes shared resources such as the database
// connection and environment variables.
type Dependency struct {
	Database *database.Database
	Env      Infrastructure.EnvVars
}

// Factory initializes and returns a new Dependency instance, creating and injecting
// the required components (e.g., database connection). It also ensures proper error
// handling if any component fails to initialize.
func Factory(envVars Infrastructure.EnvVars) (*Dependency, error) {
	database, err := database.Factory(envVars)
	if err != nil {
		return nil, err
	}

	return &Dependency{
		Database: database,
		Env:      envVars,
	}, nil
}
