package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type User struct {
	Id   int
	Name string
}

func EncodeToString(u []User) string {
	var ans string
	for i := range u {
		ans += "{name: " + u[i].Name + " id: " + strconv.Itoa(u[i].Id) + "}"
		if i+1 != len(u) {
			ans += ", "
		}
	}
	return ans
}

func nameExist(u []User, name string) bool {
	for i := range u {
		if u[i].Name == name {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8080")
	conn, _ := ln.Accept()

	var Users []User

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Command to server: ", message)

		trimMessage := strings.TrimSpace(message)
		splitMessage := strings.Split(trimMessage, " ")

		switch splitMessage[0] {
		case "register":
			if len(splitMessage) == 1 {
				conn.Write([]byte("Register without name\n"))
				log.Println("Register without name")
				break
			}
			name := strings.Replace(trimMessage, "register ", "", 1)
			if nameExist(Users, name) {
				conn.Write([]byte("This name already exists\n"))
				log.Println("User with name '" + name + "' already exists")
				break
			}
			conn.Write([]byte("register user '" + name + "' with id " + strconv.Itoa(len(Users)) + "\n"))
			log.Println("register user '" + name + "' with id: " + strconv.Itoa(len(Users)))
			Users = append(Users, User{Name: name, Id: len(Users)})

		case "list":
			conn.Write([]byte(EncodeToString(Users) + "\n"))
			log.Println("list")

		default:
			conn.Write([]byte("Incorrect message\n"))
			log.Println("Incorrect message")
		}
	}
}
