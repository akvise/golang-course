package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

const uri = "mongodb://localhost:27017/\n"

var ctx context.Context
var Database *mongo.Database
var Collection *mongo.Collection
var Client *mongo.Client

func InitDB() {
	Client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	ErrFunc(err)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(ctx)
	ErrFunc(err)
	Database = Client.Database("database")
	Collection = Database.Collection("list")
}

func AddDBRequest(id int64, message string) error{

	_, err := Collection.InsertOne(ctx, bson.D{
		{Key: strconv.FormatInt(id, 10), Value: message},
	})
	if err != nil{
		return err
	}
	return nil
}

func GetList() (string, error) {

	cur, err := Collection.Find(context.Background(), bson.D{})
	if err != nil{
		return "", err
	}
	defer cur.Close(context.Background())

	var ans string
	for cur.Next(context.Background()) {
		result := struct {
			Foo string
			Bar int32
		}{}
		err = cur.Decode(&result)
		ErrFunc(err)
		raw := cur.Current
		ans += raw.String() + "\n"

	}

	return ans, nil
}