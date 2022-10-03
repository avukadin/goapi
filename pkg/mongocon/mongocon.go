package mongocon

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/avukadin/goapi/constants"
	
	"go.mongodb.org/mongo-driver/mongo/options"	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetMongoClient() *mongo.Client{
	client, err := mongo.NewClient(options.Client().ApplyURI(constants.MONGO_URI))
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

	// Connect
	clientsCollection := client.Database("dev-api").Collection("clients")
	ctx := getMongoContext(client)	

	// Query
	cursor, err := clientsCollection.Find(ctx, bson.D{{"handle",clientHandle}})
	if err != nil {log.Fatal(err)}
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}


	if len(results) == 0 {
		return "Not Found"
	}	

	return results[0]["_id"].(primitive.ObjectID).String()
}
