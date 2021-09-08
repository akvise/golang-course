package server

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type App struct {
	collection *mongo.Collection

}

func NewApp() *App{
	db := initDb()
	list := db.Collection(viper.GetString("mongo.collection"))

	return &App{
		collection: list,
	}
}

func (app *App) Run(){


}

func initDb() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf("Error with connection to mongoDB")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	ErrFunc(err)
	database := client.Database(viper.GetString("mongo.name"))
	return database
}

func ErrFunc(err error){
	log.Fatal(err)
}
