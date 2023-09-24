package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGO_URI = "mongodb://localhost:27017"

type Top250Movie struct {
	ID          string   `bson:"_id"`
	Rank        int      `bson:"rank"`
	Name        string   `bson:"name"`
	RatingValue float64  `bson:"rating_value"`
	RatingCount int      `bson:"rating_count"`
	Quote       string   `bson:"quote"`
	Year        string   `bson:"year"`
	Genres      []string `bson:"genres"`
	RunningTime int      `bson:"running_time"`
}

func main() {
	var movie Top250Movie

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalln(err)
		}
	}()

	collection := client.Database("test").Collection("movie_top250")
	if err = collection.FindOne(ctx, bson.M{"rank": 1}).Decode(&movie); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%#v\n", movie)
}
