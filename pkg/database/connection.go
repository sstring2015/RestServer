package database

import (
	"context"
	"fmt"

	"github.com/RestServer/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB(cfg config.DatabaseConfig) (*mongo.Database, error) {
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.Username, cfg.Password, cfg.Host, cfg.Port)

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	fmt.Println("Connection established with MongoDB.")
	db := client.Database(cfg.Name)
	return db, nil
}
