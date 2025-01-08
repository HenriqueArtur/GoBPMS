package infrastructure

import (
	"os"
	"testing"
)

func createMockEnvFile(t *testing.T, content string) string {
	t.Helper()
	file, err := os.CreateTemp("", "*.env")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %s", err)
	}

	return file.Name()
}

func Test_LoadEnv_Success(t *testing.T) {
	mockContent := `
MONGO_HOST=localhost
MONGO_PORT=27017
MONGO_USER=admin
MONGO_PASSWORD=secret
MONGO_DATABASE=testdb
`
	filepath := createMockEnvFile(t, mockContent)
	defer os.Remove(filepath)

	envVars, err := LoadEnv(filepath)
	if err != nil {
		t.Fatalf("Expected no error, got: %s", err)
	}

	if envVars.DB.Host != "localhost" {
		t.Errorf("Expected MONGO_HOST to be 'localhost', got: %s", envVars.DB.Host)
	}
	if envVars.DB.Port != "27017" {
		t.Errorf("Expected MONGO_PORT to be '27017', got: %s", envVars.DB.Port)
	}
	if envVars.DB.User != "admin" {
		t.Errorf("Expected MONGO_USER to be 'admin', got: %s", envVars.DB.User)
	}
	if envVars.DB.Password != "secret" {
		t.Errorf("Expected MONGO_PASSWORD to be 'secret', got: %s", envVars.DB.Password)
	}
	if envVars.DB.Database != "testdb" {
		t.Errorf("Expected MONGO_DATABASE to be 'testdb', got: %s", envVars.DB.Database)
	}
}

func Test_LoadEnv_FileNotFound(t *testing.T) {
	_, err := LoadEnv("nonexistent.env")
	if err == nil {
		t.Fatal("Expected an error for non-existent file, got nil")
	}
}

func Test_LoadEnv_InvalidFormat(t *testing.T) {
	mockContent := `
INVALID_LINE
MONGO_HOST=localhost
MONGO_PORT
`
	filepath := createMockEnvFile(t, mockContent)
	defer os.Remove(filepath)

	_, err := LoadEnv(filepath)
	if err == nil {
		t.Fatal("Expected an error for invalid file format, got nil")
	}
}

func Test_LoadEnv_EmptyFile(t *testing.T) {
	filepath := createMockEnvFile(t, "")
	defer os.Remove(filepath)

	_, err := LoadEnv(filepath)
	if err != nil {
		t.Fatalf("Expected no error for an empty file, got: %s", err)
	}
}

func Test_LoadEnv_MissingRequiredFields(t *testing.T) {
	mockContent := `
MONGO_HOST=localhost
MONGO_PORT=27017
MONGO_USER=admin
`
	filepath := createMockEnvFile(t, mockContent)
	defer os.Remove(filepath)

	envVars, err := LoadEnv(filepath)
	if err != nil {
		t.Fatalf("Expected no error for partial config, got: %s", err)
	}

	// Verify missing fields result in default empty values
	if envVars.DB.Password != "" {
		t.Errorf("Expected MONGO_PASSWORD to be '', got: %s", envVars.DB.Password)
	}
	if envVars.DB.Database != "" {
		t.Errorf("Expected MONGO_DATABASE to be '', got: %s", envVars.DB.Database)
	}
}
