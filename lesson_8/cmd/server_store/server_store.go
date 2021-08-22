package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Token string 		`json:"token"`
	CurrentDate string	`json:"current_date"`
	ExpireAt string 	`json:"expire_at"`
}

func errFunc(err error){
	if err != nil {
		log.Fatal(err)
		return
	}
}

var Users map[string]User

func store(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		body, err := io.ReadAll(r.Body)


		res := *new(User)
		err = json.Unmarshal(body, &res)
		errFunc(err)

		Users[res.Token] = res
		log.Println(Users)

	case "GET":
		j, _ := json.MarshalIndent(Users, " ", "  ")
		fmt.Fprintf(w, string(j))
	default:
		fmt.Fprintf(w, "Incorrect method")
	}

}

func main() {
	Users = make(map[string]User)
	fmt.Println("Launching store server ...")
	http.HandleFunc("/", store)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
