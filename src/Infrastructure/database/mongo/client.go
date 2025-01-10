// Package dbmongoclient provides utilities to establish and manage connections to MongoDB.
package dbmongoclient

import (
	"context"
	"fmt"
	"time"

	infrastructure "github.com/HenriqueArtur/ProcessInGo/src/Infrastructure"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToMongoDB establishes a connection to MongoDB using the provided configuration.
// It returns a MongoDB client instance or an error if the connection fails.
func ConnectToMongoDB(config infrastructure.MongoConfig) (*mongo.Client, error) {
	// Create a context with a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB using the connection URL from the config
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URL))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Verify the connection with a ping
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return client, nil
}
