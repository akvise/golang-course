package main

import (
	"../../pkg/fib"
	"fmt"
)

func main() {
	var n int

	defer fmt.Println("Complete!")
	fmt.Print("This program can count fibonacci numbers\n", "Please enter number: ")
	fmt.Scan(&n)

	fib.NumL(n) //by loop
	fib.NumR(n)	//by recursion

}
