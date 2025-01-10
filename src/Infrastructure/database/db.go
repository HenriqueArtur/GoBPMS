// Package database provides functionality for setting up and managing database connections.
package database

import (
	infrastructure "github.com/HenriqueArtur/ProcessInGo/src/Infrastructure"
	dbmongoclient "github.com/HenriqueArtur/ProcessInGo/src/Infrastructure/database/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

// Database struct represents the database connection configuration, client, and a function to disconnect.
type Database struct {
	Env        infrastructure.MongoConfig
	Client     *mongo.Client
	Disconnect func()
}

// FactoryDatabase initializes a new Database instance by establishing a connection to Database.
func FactoryDatabase(envVars infrastructure.EnvVars) (*Database, error) {
	client, err := dbmongoclient.ConnectToMongoDB(envVars.DB)
	if err != nil {
		return nil, err
	}

	disconnect := func() {
		client.Disconnect(nil)
	}

	return &Database{
		Env:        envVars.DB,
		Client:     client,
		Disconnect: disconnect,
	}, nil
}
