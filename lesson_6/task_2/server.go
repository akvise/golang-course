package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8080")
	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", message)
		r := []rune(message)
		r = r[:len(r)-2]

		var newMessage string
		i, err := strconv.Atoi(string(r))
		if err == nil {
			newMessage = strconv.Itoa(i * 2)
		} else {
			newMessage = strings.ToUpper(message)
		}
		conn.Write([]byte(newMessage + "\n"))
	}
}
