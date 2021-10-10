package config

import (
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() (*mongo.Database, error) {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, err := strconv.Atoi(CONFIG["MONGO_POOL_MIN"])
	if err != nil {
		return nil, err
	}

	mongoPoolMax, err := strconv.Atoi(CONFIG["MONGO_POOL_MAX"])
	if err != nil {
		return nil, err
	}

	mongoMaxIdleTime, err := strconv.Atoi(CONFIG["MONGO_MAX_IDLE_TIME_SECOND"])
	if err != nil {
		return nil, err
	}

	option := options.Client().
		ApplyURI(CONFIG["MONGO_URI"]).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	database := client.Database(CONFIG["MONGO_DATABASE"])
	return database, nil
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
