package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"../User"
)

func PostToStore(token string, name string, address string) {
	curUser := User.Struct{Token: token, CurrentDate: time.Now().String(),
		ExpireAt: time.Now().AddDate(0, 0, 10).String()}
	mapUser := map[string]User.Struct{name + " " + address: curUser}
	json_data, err := json.Marshal(mapUser)
	if err != nil {
		log.Println(err, token)
	}

	r, err := http.Post("http://localhost:8081/", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Println(err, token)
	}
	fmt.Println(r)

}

func root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 root not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "index.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		name := r.FormValue("name")
		address := r.FormValue("address")

		cookie := &http.Cookie{
			Name:  "token",
			Value: name + ":" + address,
		}

		http.SetCookie(w, cookie)
		PostToStore(cookie.Value, name, address)

		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
		fmt.Fprintf(w, "Token = %s", cookie.Value)

		fmt.Println("Post request: ")
		fmt.Printf("Name = %s\nAddress = %s\nToken = %s\n\n", name, address, cookie.Value)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	fmt.Println("Launching server ...")
	http.HandleFunc("/", root)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
