package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

//ConnectMongo function to establish mongo connection
func ConnectMongo() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@mongo:27017"))
	if err != nil {
		log.Panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Panic(err)
	}
	// defer client.Disconnect(ctx)
	return client
}

//GetSharedConnection get db connection
func GetSharedConnection() *mongo.Database {
	return ConnectMongo().Database("sellerapp")
}
