package database

import (
	"context"
	"encoding/json"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"main/entity"
	"time"
)

type Repository struct {
	Database   *mongo.Database
	Collection *mongo.Collection
	Context 	context.Context
}

func InitDb() *Repository {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf("Error with connection to mongoDB")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
	}
	database := client.Database(viper.GetString("mongo.name"))
	collection := database.Collection(viper.GetString("mongo.collection"))

	return &Repository{
		Database: database,
		Collection: collection,
		Context: ctx,
	}
}

func Set(repository Repository,response entity.Response) error{
	_, err := repository.Collection.InsertOne(repository.Context, response)
	if err != nil{
		return err
	}
	return nil
}

func Get(repository Repository) (string, error) {
	cur, err := repository.Collection.Find(context.Background(), bson.D{})
	if err != nil{
		return "", err
	}
	defer cur.Close(context.Background())

	var message string
	for cur.Next(context.Background()) {
		var result entity.Response
		err = cur.Decode(&result)
		if err != nil{
			log.Println(err)
		}
		raw := cur.Current
		var m map[string]interface{}
		json.Unmarshal([]byte(raw.String()), &m)
		delete(m, "_id")

		j, _ := json.Marshal(m)
		message += string(j) + "\n"

	}

	return message, nil
}