package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()
	fmt.Println("You can use 2 command:\n" +
		"1. list\t\t- return list of users that store on the server\n" +
		"2. register <name> - append new user with the name you indicated")

	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print("server: " + message)
	}
}
