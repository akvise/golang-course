package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func (m Musician) String() string {
	j, _ := json.MarshalIndent(m, "", "  ")
	return string(j)
}

type Musician struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

func Parse(s string) string {
	return strings.ReplaceAll(s, "call-me-", "")
}

func main() {
	client := http.Client{}
	resp, err := client.Get("http://localhost:8080/api/v0/user")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//encode slice byte to struct
	Musicians := make([]Musician, 0)
	err = json.Unmarshal(body, &Musicians)

	for {
		fmt.Print("Number: ")

		reader := bufio.NewReader(os.Stdin)
		num, _ := reader.ReadString('\n')
		num = strings.TrimSpace(num)

		if num == "exit" {
			break
		}

		correctNum := false
		for i := range Musicians {
			m := Musicians[i]
			if num == Parse(m.Phone) {
				fmt.Println(m)
				correctNum = true
			}
		}
		if correctNum == false {
			fmt.Println("Undefined number. Please, try again.")
		}
	}
}
