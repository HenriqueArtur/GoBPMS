// Package infrastructure provides utilities for loading environment variables
// and structuring configuration for various parts of the application.
package infrastructure

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// EnvVars represents the overall application environment configuration.
// It includes MongoDB configuration and can be extended for additional services.
type EnvVars struct {
	DB MongoConfig
}

// MongoConfig holds MongoDB-specific configurations.
// This structure includes details like the connection URL, host, port, user, password, and database name.
type MongoConfig struct {
	URL      string
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// LoadEnv loads environment variables from the specified file and returns an EnvVars object.
// This function reads the provided `.env` file (specified by filepath), parses it, and organizes the
// values into the EnvVars structure. If the file cannot be read or contains invalid entries, an error is returned.
func LoadEnv(filepath string) (*EnvVars, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open environment file '%s': %w", filepath, err)
	}
	defer file.Close()

	envVars, err := mapEnvs(file, filepath)
	if err != nil {
		return nil, err
	}

	// Return the EnvVars object
	return &EnvVars{
		DB: mountDbObj(envVars),
	}, nil
}

func mapEnvs(file *os.File, filepath string) (map[string]string, error) {
	envVars := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid entry in '%s': %s", filepath, line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		envVars[key] = value
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading environment file '%s': %w", filepath, err)
	}
	return envVars, nil
}

func mountDbObj(envVars map[string]string) MongoConfig {
	host := envVars["MONGO_HOST"]
	port := envVars["MONGO_PORT"]
	user := envVars["MONGO_USER"]
	password := envVars["MONGO_PASSWORD"]
	database := envVars["MONGO_DATABASE"]

	url := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", user, password, host, port, database)

	return MongoConfig{
		URL:      url,
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}
}
