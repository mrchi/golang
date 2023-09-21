package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Top250Movie struct {
	ID          string   `bson:"_id"`
	Rank        int      `bson:"rank"`
	Name        string   `bson:"name"`
	RatingValue float32  `bson:"rating_value"`
	RatingCount int      `bson:"rating_count"`
	Quote       string   `bson:"quote"`
	Year        string   `bson:"year"`
	Genres      []string `bson:"genres"`
	RunningTime int      `bson:"running_time"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln(err)
	}
}
