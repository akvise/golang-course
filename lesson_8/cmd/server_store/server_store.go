package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../User"
)

var Users []map[string]User.Struct

func store(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("something wrong")
		return
	}

	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		var res map[string]User.Struct
		json.NewDecoder(r.Body).Decode(&res)
		Users = append(Users, res)

	case "GET":
		j, _ := json.MarshalIndent(Users, "", "  ")
		fmt.Fprintf(w, string(j))
	default:
		fmt.Fprintf(w, "Incorrect method")
	}

}

func main() {
	fmt.Println("Launching store server ...")
	http.HandleFunc("/", store)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
