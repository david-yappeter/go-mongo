package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ConnectMongo Connect Mongo
func ConnectMongo() *mongo.Client {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseConfig := fmt.Sprintf("mongodb+srv://%s:%s@op-op-club.gwjgw.mongodb.net/%s?retryWrites=true&w=majority", os.Getenv("USER_MONGO"), os.Getenv("PASS_MONGO"), os.Getenv("DB_MONGO"))

	clientOptions := options.Client().ApplyURI(databaseConfig)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	//Check Connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
