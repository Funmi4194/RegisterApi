package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//create a connect function that returns a client if successfully connected
func ConnectDB() *mongo.Client {
	//creates clientOptions and connect it to the databases
	clientOptions := options.Client().ApplyURI("mongodb+srv://funmi-tech:funmi4194@cluster0.v5gihly.mongodb.net/?retryWrites=true&w=majority")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client

}

//pointer to a mongo.client(client inatance)
var DB *mongo.Client = ConnectDB() //DB can be used to create collection instances in the code base

//create a function that get collection name
func GetConnection(collectionInstance string) *mongo.Collection {
	db := DB.Database("RegisterAPI")
	collection := db.Collection(collectionInstance)
	return collection
}
