package mongoclient

import (
	"testing"

	infrastructure "github.com/HenriqueArtur/ProcessInGo/src/Infrastructure"
)

func Test_ConnectToMongoDB_Success(t *testing.T) {
	// Example MongoConfig for a local MongoDB server
	config := infrastructure.MongoConfig{
		URL:      "mongodb://localhost:27017",
		Host:     "localhost",
		Port:     "27017",
		User:     "root",
		Password: "changeThePassword",
		Database: "",
	}

	client, err := Connect(config)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %s", err)
	}
	defer client.Disconnect(nil)
}

func Test_ConnectToMongoDB_InvalidURL(t *testing.T) {
	// Example of an invalid MongoConfig
	config := infrastructure.MongoConfig{
		URL: "invalid_url",
	}

	_, err := Connect(config)
	if err == nil {
		t.Fatal("Expected error for invalid MongoDB URL, got nil")
	}
}
