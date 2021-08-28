package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func AddDBRequest(city string, value string) error{

	_, err := Collection.InsertOne(ctx, bson.D{
		{Key: "city", Value: city},
		{Key: "value", Value: value},
		{Key: "time_requested", Value: time.Now().String()},
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
		var m map[string]interface{}
		json.Unmarshal([]byte(raw.String()), &m)
		delete(m, "_id")

		j, _ := json.Marshal(m)
		ans += string(j) + "\n"

	}

	return ans, nil
}