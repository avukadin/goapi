package mongocon

import (
	"context"
	"log"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	
	"go.mongodb.org/mongo-driver/mongo/options"	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetMongoClient() *mongo.Client{
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func getMongoContext(client *mongo.Client) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return ctx
}


func GetClientID(client *mongo.Client, clientHandle string) string{
	clientsCollection := client.Database("api").Collection("users")
	ctx:=getMongoContext(client)	
	
	cursor, err := clientsCollection.Find(ctx, bson.D{{"email",clientHandle}})

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
        fmt.Println(result["_id"])
	}
	return results[0]["_id"].(primitive.ObjectID).String()
}
