package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

	// Insert BSON document
	doc := categorie{"Aji", "Tibco Team"}

	insertResult, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	defer client.Disconnect(context.TODO())

	// update BSON document
	filter := bson.D{{"name", "Aji"}}

	update := bson.D{
		{"$set", bson.D{
			{"desc", "Tibco Support"},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// delete BSON document

	// deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{"desc", "Tibco Team"}})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	// select / find BSON document
	var results []*categorie
	// 1 --
	findOptions := options.Find()
	//findOptions.SetLimit(2)

	// 2 -- Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 3 --
	for cur.Next(context.TODO()) {
		var record categorie
		err := cur.Decode(&record)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &record)
	}

	for _, v := range results {
		// fmt.Printf("%v -- %v", v.Name, v.Desc)
		fmt.Println("%v -- %v", v.Name, v.Desc)
	}

}
