package main

import (
	"../../pkg/fibonacci"
	"fmt"
)

func main() {
	var n int

	defer fmt.Println("\nCompleted!")
	fmt.Print("This program can count fibonacci numbers\n", "Please enter number: ")
	fmt.Scan(&n)


	fmt.Println("Count by with liner execution time:")
	for i := 1; i < n; i++ {
		fmt.Print(fibonacci.FibLoop(i)," ") 	//by loop
	}

	fmt.Println("\nCount by recursion with exponential execution time:")
	for i := 1; i < n; i++ {
		fmt.Print(fibonacci.FibRecursion(i), " ") //by recursion
	}
}
