package mongo

import (
	"context"
	"log/slog"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectTimeout = 5

type DB struct {
    client *mongo.Client
    database *mongo.Database
    logger *slog.Logger
}

func Initialize(logger *slog.Logger) *DB {
    uri := os.Getenv("DATABASE_URL")
    name := os.Getenv("DATABASE_NAME")
    result := new(DB)
    result.logger = logger
    result.client, _ = result.connect(uri)
    result.database = result.client.Database(name)
    return result
}

func (db *DB) connect(uri string) (*mongo.Client, context.CancelFunc) {
    ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
    options := options.Client().ApplyURI(uri).SetServerSelectionTimeout(time.Second * 60)
    client, err := mongo.Connect(ctx, options)
    if err != nil {
        db.logger.Error("connecting to database", "service", "mongodb", "error", err.Error())
        os.Exit(1)
    }
    if err = client.Ping(ctx, nil); err != nil {
        db.logger.Error("pinging database", "service", "mongodb", "error", err.Error())
        os.Exit(1)
    }
    db.logger.Info("initialized", "service", "mongodb")
    return client, cancel
}

func (db *DB) Database() *mongo.Database {
    return db.database
}

func (db *DB) Close(ctx context.Context) error {
    return db.client.Disconnect(ctx)
}
