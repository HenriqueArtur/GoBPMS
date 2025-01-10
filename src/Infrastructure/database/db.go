// Package database provides functionality for setting up and managing database connections.
package database

import (
	Infrastructure "github.com/HenriqueArtur/ProcessInGo/src/Infrastructure"
	MongoClient "github.com/HenriqueArtur/ProcessInGo/src/Infrastructure/database/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

// Database struct represents the database connection configuration, client, and a function to disconnect.
type Database struct {
	Env    Infrastructure.MongoConfig
	Client *mongo.Client
}

// Factory initializes a new Database instance by establishing a connection to Database.
func Factory(envVars Infrastructure.EnvVars) (*Database, error) {
	client, err := MongoClient.Connect(envVars.DB)
	if err != nil {
		return nil, err
	}

	return &Database{
		Env:    envVars.DB,
		Client: client,
	}, nil
}

// Disconnect from database
func (d Database) Disconnect() {
	d.Client.Disconnect(nil)
}
