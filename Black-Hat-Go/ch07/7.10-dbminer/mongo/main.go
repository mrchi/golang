package main

import (
	"context"
	"errors"
	"log"

	"github.com/mrchi/golang/Black-Hat-Go/ch07/7.10-dbminer/dbminer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoMiner struct {
	URI    string
	Client *mongo.Client
}

func (m *MongoMiner) connect() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.URI))
	if err != nil {
		return err
	}
	m.Client = client
	return nil
}

func New(uri string) (*MongoMiner, error) {
	m := MongoMiner{URI: uri}
	if err := m.connect(); err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *MongoMiner) GetSchema() (*dbminer.Schema, error) {
	var s = new(dbminer.Schema)

	dbNames, err := m.Client.ListDatabaseNames(context.TODO(), &options.ListDatabasesOptions{})
	if err != nil {
		return nil, err
	}

	for _, dbName := range dbNames {
		db := dbminer.Database{Name: dbName}
		collectionNames, err := m.Client.Database(dbName).ListCollectionNames(context.TODO(), &options.ListCollectionsOptions{})
		if err != nil {
			return nil, err
		}

		for _, collectionName := range collectionNames {
			table := dbminer.Table{Name: collectionName}
			var docRaw bson.D
			err := m.Client.Database(dbName).Collection(collectionName).FindOne(context.TODO(), bson.M{}, &options.FindOneOptions{}).Decode(&docRaw)
			if err != nil {
				// 如果 collection 为空，则跳过当前 collection 继续循环
				if errors.Is(err, mongo.ErrNoDocuments) {
					continue
				} else {
					return nil, err
				}
			}
			for _, f := range docRaw {
				table.Columns = append(table.Columns, f.Key)
			}
			db.Tables = append(db.Tables, table)
		}
		s.Databases = append(s.Databases, db)
	}
	return s, nil
}

func main() {
	miner, err := New("mongodb://localhost:27017")
	if err != nil {
		log.Panicln(err)
	}

	if err := dbminer.Search(miner); err != nil {
		log.Panicln(err)
	}
}
