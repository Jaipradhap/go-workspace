package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type categorie struct {
	Name string
	Desc string
}

func main() {
	fmt.Println("Mongo DB connects..")

	// session, err := mgo.Dail("localhost")
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.close()

	// session.SetMode(mgo.Monotonic, true)

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("taskdb").Collection("categories")

	doc := categorie{"Aji", "Tibco Team"}

	insertResult, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	defer client.Disconnect(context.TODO())
}
