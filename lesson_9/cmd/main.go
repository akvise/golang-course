package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Phone string
	Group int
}

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Person{})

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
	var persons []Person
	db.Find(&persons)
	json.NewEncoder(w).Encode(&persons)
}


func AddPersonPost(w http.ResponseWriter, r *http.Request){
	var person Person
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	person.Name = r.FormValue("name")
	person.Phone = r.FormValue("phone")
	person.Group = len([]rune(person.Name))

	db.Create(&person)

	fmt.Fprintf(w, "Name = %s\n", person.Name)
	fmt.Fprintf(w, "Phone = %s\n", person.Phone)
	fmt.Fprintf(w, "Group = %d\n", person.Group)

	fmt.Printf("Name = %s\nAddress = %s\nGroup = %d\n\n", person.Name, person.Phone, person.Group)
}

func AddPersonGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}