package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance

func ConnectDB() error {

	err := godotenv.Load()

	if err != nil {
		return errors.New("FAILED TO LOAD ENV")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		return err
	}

	fmt.Println("Connected to database")

	MI = MongoInstance{
		Client: client,
		DB:     client.Database(os.Getenv("MONGO_DB")),
	}
	return nil
}
