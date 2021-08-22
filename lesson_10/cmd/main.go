package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"strconv"
	"time"
)

const uri = "mongodb://localhost:27017/\n"

var ctx context.Context
var Database *mongo.Database
var Collection *mongo.Collection

func main() {
	// init database
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	errFunc(err)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	errFunc(err)
	defer client.Disconnect(ctx)
	Database = client.Database("database")
	Collection = Database.Collection("book")


	// launching server and handle requests
	router := mux.NewRouter()
	router.HandleFunc("/", root).Methods("GET")
	router.HandleFunc("/list/", GetList).Methods("GET")
	router.HandleFunc("/add/", AddPersonGet).Methods("GET")
	router.HandleFunc("/add/", AddPersonPost).Methods("POST")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))

}

func root(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<a href=http://localhost:8080/list/> LIST </a> <br><br>" +
		"<a href=http://localhost:8080/add>ADD</a>")
}

func GetList(w http.ResponseWriter, r *http.Request) {
	cur, err := Collection.Find(context.Background(), bson.D{})
	errFunc(err)
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		result := struct{
			Foo string
			Bar int32
		}{}
		err = cur.Decode(&result)
		errFunc(err)
		raw := cur.Current

		fmt.Fprintf(w, raw.String() + "\n")
	}
}


func AddPersonPost(w http.ResponseWriter, r *http.Request){
	errFunc(r.ParseForm())

	Name := string(r.FormValue("name"))
	Phone := string(r.FormValue("phone"))
	Group := len([]rune(Name))

	// save data in mongo
	_, err := Collection.InsertOne(ctx, bson.D{
		{Key: Name, Value: bson.A{Phone, strconv.Itoa(Group)}},
	})
	errFunc(err)

	fmt.Fprintf(w,"Name = %s\nAddress = %s\nGroup = %d\n\n", Name, Phone, Group)
	fmt.Printf("Name = %s\nAddress = %s\nGroup = %d\n\n", Name, Phone, Group)
}

func AddPersonGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func errFunc(err error){
	if err != nil { log.Println(err) }
}