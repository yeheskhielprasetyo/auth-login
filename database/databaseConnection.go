package database

import (
	"context"
	"fmt"
	"go/mongodb.org/mongo-driver/mongo/options"
	"log" // log out error
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("tidak terhubung ke env")
	}

	MongoDb := os.Getenv("MONGODB_URL")

	mongo.NewClient(options.Client().ApplyURI(MongoDb))

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("koneksi ke db")

	return client
}

var Client *mongo.client = DBinstance()

// func akses collection db

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}
