package main

import (
	"fmt"
	"log"
	"net/http"
)

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
			Value: name+":"+address,
		}
		http.SetCookie(w, cookie)

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
	fmt.Println("Launching server (port:80)...")
	http.HandleFunc("/", root)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
