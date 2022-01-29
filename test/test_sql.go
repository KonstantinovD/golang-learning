package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("testdb").
		Collection("stations")
}

type Station struct {
	ID      primitive.ObjectID `bson:"_id"`
	LocalId int32              `bson:"local_id"`
	Name    string             `bson:"name"`
	Lat     float64            `bson:"lat"`
	Lng     float64            `bson:"lng"`
}

func main() {
	filter := bson.D{{}}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return
	}
	var stations []*Station

	for cur.Next(ctx) {
		var st Station
		err := cur.Decode(&st)
		if err != nil {
			log.Fatal(err)
		}
		stations = append(stations, &st)
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	for _, o := range stations {
		fmt.Println(o)
	}
}
