package fib

import "fmt"

// NumL : count Fibonacci number by loop and print
func NumL(n int) {
	cur, prev :=  1, 0
	for i := 0; i < n - 1; i++ {
		prev, cur = cur, cur + prev
	}
	fmt.Println(cur)
}

// rec : count Fibonacci number by recursion
func rec(n int) int {
	if n == 1 || n == 2{
		return 1
	}
	return rec(n - 1) + rec(n - 2)
}

// NumR : print recursion func "rec" with Fib num
func NumR(n int){
	fmt.Println(rec(n))
}
