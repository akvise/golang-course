package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Host       string              `json:"host: "`
	UserAgent  string              `json:"user_agent: "`
	RequestUri string              `json:"request_uri: "`
	Headers    map[string][]string `json:"headers: "`
}

func (u User) String() string {
	j, _ := json.MarshalIndent(u, "", "  ")
	return string(j)
}

func currentUser(r *http.Request) User {
	var u User
	u.Headers = make(map[string][]string)
	u.Host = r.Host
	u.UserAgent = r.UserAgent()
	u.RequestUri = r.RequestURI
	u.Headers["Accept"] = r.Header["Accept"]
	u.Headers["User-Agent"] = r.Header["User-Agent"]

	return u
}

func root(w http.ResponseWriter, r *http.Request) {
	u := currentUser(r)
	fmt.Fprintf(w, u.String())
}

func main() {
	fmt.Println("Start to port 8080...")
	http.HandleFunc("/", root)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
